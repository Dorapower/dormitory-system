package model

import (
	"dormitory-system/src/cache"
	"dormitory-system/src/database"
	"sync"
	"time"
)

type Orders struct {
	ID            int `gorm:"primaryKey;autoIncrement"`
	Uid           int
	GroupId       int `gorm:"default:0"`
	BuildingId    int
	SubmitTime    int
	CreateTime    int
	FinishTime    int `gorm:"default:0"`
	RoomId        int
	ResultContent string
	Remarks       string
	IsDel         int `gorm:"default:0"`
	Status        int `gorm:"default:0"`
}

// CreateOrder : creat group's or personal order
func CreateOrder(uid, groupId, buildingId, submitTime int) int {
	var db = database.MysqlDb
	var orderId int
	var gender int

	db.Model(Users{}).Select("gender").Where("uid = ?", uid).Scan(&gender)
	bedCount := cache.GetBuildingCache(buildingId, gender)

	var stuCnt int64 // group members count
	db.Model(GroupsUser{}).Select("uid").Where("group_id = ? and is_del = ?", groupId, 0).Count(&stuCnt)

	if groupId == 0 && bedCount > 0 {
		//TODO: to queue

		orderId = DealPersonalOrder(uid, buildingId, submitTime)
	} else if groupId != 0 && bedCount >= int(stuCnt) {
		//TODO: to queue

		orderId = DealGroupOrder(uid, groupId, buildingId, submitTime)
	} else { // no bed to choose
		var order Orders
		order.CreateTime = int(time.Now().Unix())
		order.Uid = uid
		order.SubmitTime = submitTime
		order.BuildingId = buildingId
		order.Remarks = "none"
		order.RoomId = 0
		order.ResultContent = "no available room"
		order.Status = 1
		order.FinishTime = int(time.Now().Unix())
		order.GroupId = groupId
		if db.Create(&order).Error != nil {
			return -1
		}
		return order.ID
	}
	return orderId
}
func MatchUserGroup(uid, groupId int) bool {
	var db = database.MysqlDb
	var groupids []int
	db.Model(GroupsUser{}).Select("group_id").Where("uid = ? and is_del = ?", uid, 0).Scan(&groupids)
	for _, id := range groupids {
		if id == groupId {
			return true
		}
		if groupId == 0 { // is in a group but send request to create personal order
			return false
		}
	}
	if groupId == 0 { // is not in a group and send request to create personal order
		return true
	}
	return false
}
func DealPersonalOrder(uid, buildingId, submitTime int) int {
	var db = database.MysqlDb
	var order Orders
	order.CreateTime = int(time.Now().Unix())
	order.Uid = uid
	order.SubmitTime = submitTime
	order.BuildingId = buildingId
	order.Remarks = "none"
	order.RoomId = 0

	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	//check if user have dormitory
	var tempBed Beds
	db.Model(Beds{}).Where("uid = ? and is_valid = ? and is_del = ? and status = ?", uid, 1, 0, 1).First(&tempBed)
	if tempBed.ID != 0 {
		order.ResultContent = "already have bed"
		order.Status = 2
		order.FinishTime = int(time.Now().Unix())
		if db.Create(&order).Error != nil {
			return -1
		}
		return order.ID
	}

	// get user's gender
	var gender int
	db.Model(Users{}).Select("gender").Where("uid = ?", uid).Scan(&gender)

	// get available room
	var roomIds []int
	db.Model(Rooms{}).Select("id").Where("building_id = ? and gender = ? and is_valid = ? and is_del = ?", buildingId, gender, 1, 0).Order("order_num").Scan(&roomIds)

	// select the first matched bed from beds ordered by "order_num"
	for _, roomId := range roomIds {
		var bed Beds
		result := db.Model(Beds{}).Where("room_id = ? and is_valid = ? and is_del = ? and status = ?", roomId, 1, 0, 0).Order("order_num").First(&bed)

		// find it
		if result.Error == nil {
			order.RoomId = bed.RoomId
			var roomName string
			db.Model(Rooms{}).Select("Name").Where("id = ?", bed.RoomId).Scan(&roomName)
			order.ResultContent = "success: " + roomName
			// update bed's information
			db.Model(&bed).Updates(map[string]interface{}{"uid": uid, "status": 1})

			// update cache
			roomCache := cache.GetRoomCache(roomId)
			_ = cache.SetRoomCache(roomCache, roomCache-1)
			buildingCache := cache.GetBuildingCache(buildingId, gender)
			_ = cache.SetBuildingCache(buildingId, gender, buildingCache-1)

			break
		}
	}

	// fail
	if order.RoomId == 0 {
		order.ResultContent = "no available room"
	}

	order.Status = 1
	order.FinishTime = int(time.Now().Unix())
	if db.Create(&order).Error != nil {
		return -1
	}
	return order.ID
}

func DealGroupOrder(uid, groupId, buildingId, submitTime int) int {
	var db = database.MysqlDb
	var order Orders
	order.CreateTime = int(time.Now().Unix())
	order.SubmitTime = submitTime
	order.BuildingId = buildingId
	order.GroupId = groupId
	order.Remarks = "none"
	order.Uid = uid
	order.RoomId = 0

	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	//check if group have dormitory
	tempStatus := 0
	db.Model(Groups{}).Select("status").Where("id = ? and is_del = ?", groupId, 0).Scan(&tempStatus)
	if tempStatus == 1 {
		order.ResultContent = "already have room"
		order.Status = 2
		order.FinishTime = int(time.Now().Unix())
		if db.Create(&order).Error != nil {
			return -1
		}
		return order.ID
	}

	var stuCnt int64 // group members count
	var uIds []int
	db.Model(GroupsUser{}).Select("uid").Where("group_id = ? and is_del = ?", groupId, 0).Count(&stuCnt).Scan(&uIds)

	// get group gender
	var gender int
	db.Model(Users{}).Select("gender").Where("uid = ?", uid).Scan(&gender)

	// get all optional rooms
	var roomIds []int
	db.Model(Rooms{}).Select("id").Where("building_id = ? and gender = ? and is_valid = ? and is_del = ?", buildingId, gender, 1, 0).Order("order_num").Scan(&roomIds)

	for _, roomId := range roomIds {
		var bedCnt int // room's available beds
		bedCnt = cache.GetRoomCache(roomId)

		// if beds number >= people number
		if bedCnt >= int(stuCnt) {
			var bedIds []int
			db.Model(Beds{}).Limit(int(stuCnt)).Select("id").Where("room_id = ? and is_valid = ? and is_del = ? and status = ?", roomId, 1, 0, 0).Order("order_num").Scan(&bedIds)
			order.RoomId = roomId
			var roomName string
			db.Model(Rooms{}).Select("Name").Where("id = ?", roomId).Scan(&roomName)
			order.ResultContent = "success: " + roomName

			// update group information
			db.Model(Groups{}).Where("id = ?", groupId).Update("status", 1)

			// update beds information
			for i := 0; i < int(stuCnt); i++ {
				db.Model(Beds{}).Where("id = ?", bedIds[i]).Updates(map[string]interface{}{"uid": uIds[i], "status": 1})
			}

			// update cache
			roomCache := cache.GetRoomCache(roomId)
			_ = cache.SetRoomCache(roomCache, roomCache-1)
			buildingCache := cache.GetBuildingCache(buildingId, gender)
			_ = cache.SetBuildingCache(buildingId, gender, buildingCache-1)
			break
		}
	}

	// fail
	if order.RoomId == 0 {
		order.ResultContent = "no available room"
		order.Status = 2
	} else {
		order.Status = 1
	}

	order.FinishTime = int(time.Now().Unix())
	if db.Create(&order).Error != nil {
		return -1
	}
	return order.ID
}

type OrderListApi struct {
	BuildingName  string `json:"building_name"`
	GroupName     string `json:"group_name"`
	OrderId       int    `json:"order_id"`
	ResultContent string `json:"result_content"`
	Status        int    `json:"status"`
	SubmitTime    string `json:"submit_time"`
}

func GetOrderList(uid int) (orderLA []OrderListApi) {
	var db = database.MysqlDb
	var orderIds []int
	db.Model(Orders{}).Select("id").Where("uid = ?", uid).Scan(&orderIds)
	for _, id := range orderIds {
		var order Orders
		db.Where("id = ?", id).First(&order)
		var groupName string
		db.Model(Groups{}).Select("Name").Where("id = ?", order.GroupId).Scan(&groupName)
		var buildingName string
		db.Model(Buildings{}).Select("Name").Where("id = ?", order.BuildingId).Scan(&buildingName)
		submitTime := time.Unix(int64(order.SubmitTime), 0).Format("2006-01-02 15:04:05")
		var orderApi = OrderListApi{
			OrderId:       id,
			GroupName:     groupName,
			BuildingName:  buildingName,
			SubmitTime:    submitTime,
			ResultContent: order.ResultContent,
			Status:        order.Status,
		}
		orderLA = append(orderLA, orderApi)
	}
	return
}

type OrderInfoApi struct {
	RoomID int `json:"room_id"`
	Status int `json:"status"`
}

func GetOrderInfo(orderId int) (oIA OrderInfoApi) {
	var db = database.MysqlDb
	db.Model(Orders{}).Select("status", "room_id").Where("id = ?", orderId).Scan(&oIA)
	return
}
