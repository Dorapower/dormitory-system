package model

import (
	"dormitory-system/src/database"
	"time"
)

type Orders struct {
	ID            int `gorm:"primaryKey;autoIncrement"`
	Uid           int
	GroupId       int `gorm:"default:0"`
	BuildingId    int
	SubmitTime    int
	CreatTime     int
	FinishTime    int `gorm:"default:0"`
	RoomId        int
	ResultContent string
	Remarks       string
	IsDel         int `gorm:"default:0"`
	Status        int `gorm:"default:0"`
}

// CreateOrder : creat group's or personal order
func CreateOrder(uid, groupId, buildingId, submitTime int) int {
	var orderId int
	if groupId == 0 {
		orderId = DealPersonalOrder(uid, buildingId, submitTime)
	}
	orderId = DealGroupOrder(uid, groupId, buildingId, submitTime)
	return orderId
}

func DealPersonalOrder(uid, buildingId, submitTime int) int {
	var db = database.MysqlDb
	var order Orders
	order.CreatTime = int(time.Now().Unix())
	order.Uid = uid
	order.SubmitTime = submitTime
	order.BuildingId = buildingId
	order.Remarks = "none"
	order.RoomId = 0

	//check if user have dormitory
	var tempBed Beds
	result := db.Model(Beds{}).Where("uid = ? and is_valid = ? and is_del = ? and status = ?", uid, 1, 0, 1).First(&tempBed)
	if result.Error == nil {
		order.ResultContent = "already have bed"
		order.Status = 2
		order.FinishTime = int(time.Now().Unix())
		db.Create(&order)
		return order.ID
	}

	// get user's gender
	var gender int
	db.Model(Users{}).Select("gender").Where("uid = ?", uid).Scan(&gender)

	// get available room
	var roomIds []int
	db.Model(Rooms{}).Select("id").Where("building_id = ? and gender = ? and is_valid = ? and is_del = ?", buildingId, gender, 1, 0).Order("order_num").Scan(&roomIds)

	// select the first matched bed from beds ordered by "order_num"
	for roomId := range roomIds {
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
			break
		}
	}

	// fail
	if order.RoomId == 0 {
		order.ResultContent = "no available room"
	}

	order.Status = 1
	order.FinishTime = int(time.Now().Unix())
	db.Create(&order)
	return order.ID
}

func DealGroupOrder(uid, groupId, buildingId, submitTime int) int {
	var db = database.MysqlDb
	var order Orders
	order.CreatTime = int(time.Now().Unix())
	order.SubmitTime = submitTime
	order.BuildingId = buildingId
	order.GroupId = groupId
	order.Remarks = "none"
	order.Uid = uid
	order.RoomId = 0

	//check if group have dormitory
	tempStatus := 0
	db.Model(Groups{}).Select("status").Where("id = ? and is_del = ?", groupId, 0).Scan(&tempStatus)
	if tempStatus == 1 {
		order.ResultContent = "already have bed"
		order.Status = 2
		order.FinishTime = int(time.Now().Unix())
		db.Create(&order)
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

	for roomId := range roomIds {
		var bedCnt int64 // room's optional beds
		db.Model(Beds{}).Where("room_id = ? and is_valid = ? and is_del = ? and status = ?", roomIds, 1, 0, 0).Count(&bedCnt)

		// if beds number >= people number
		if bedCnt >= stuCnt {
			var bedIds []int
			db.Model(Beds{}).Limit(int(stuCnt)).Select("id").Where("room_id = ? and is_valid = ? and is_del = ? and status = ?", roomIds, 1, 0, 0).Order("order_num").Scan(&bedIds)
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
			break
		}
	}

	// fail
	if order.RoomId == 0 {
		order.ResultContent = "no available room"
	}

	order.Status = 1
	order.FinishTime = int(time.Now().Unix())
	db.Create(&order)
	return order.ID
}

type OrderListApi struct {
	order_id       int
	group_name     string
	building_name  string
	submit_time    string
	result_content string
	status         int
}

func GetOrderList(uid int) (orderLA []OrderListApi) {
	var db = database.MysqlDb
	var orderIds []int
	db.Model(Orders{}).Select("id").Where("uid = ?", uid).Scan(&orderIds)
	for id := range orderIds {
		var order Orders
		db.Where("id = ?", id).First(&order)
		var groupName string
		db.Model(Groups{}).Select("Name").Where("id = ?", order.GroupId).Scan(&groupName)
		var buildingName string
		db.Model(Buildings{}).Select("Name").Where("id = ?", order.BuildingId).Scan(&buildingName)
		submitTime := time.Unix(int64(order.SubmitTime), 0).Format("2006-01-02 15:04:05")
		var orderApi = OrderListApi{
			order_id:       id,
			group_name:     groupName,
			building_name:  buildingName,
			submit_time:    submitTime,
			result_content: order.ResultContent,
			status:         order.Status,
		}
		orderLA = append(orderLA, orderApi)
	}
	return
}

type OrderInfoApi struct {
	status  int
	room_id int
}

func GetOrderInfo(orderId int) (oIA OrderInfoApi) {
	var db = database.MysqlDb
	db.Model(Orders{}).Select("status", "room_id").Where("id = ?", orderId).Scan(&oIA)
	return
}
