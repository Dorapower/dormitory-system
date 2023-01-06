package model

import (
	"dormitory-system/src/cache"
	"dormitory-system/src/database"
)

type Beds struct {
	ID       int `gorm:"primaryKey;autoIncrement"`
	Uid      int `gorm:"default:0"`
	RoomId   int
	Name     string
	OrderNum int
	IsValid  int `gorm:"default:1"`
	Remarks  string
	IsDel    int `gorm:"default:0"`
	Status   int `gorm:"default:0"`
}

// GetMyRoomByUid : get my room's Name and roommates' Name
func GetMyRoomByUid(uid int) (roomId int, roomName string, roomMate []map[string]string) {
	var db = database.MysqlDb

	// get room's id
	db.Model(Beds{}).Select("room_id").Where("uid = ? and is_valid = ? and is_del = ? and status = ?", uid, 1, 0, 1).Scan(&roomId)

	// get room's Name
	db.Model(Rooms{}).Select("Name").Where("id = ?", roomId).Scan(&roomName)

	// get roommate's Name
	var names []string
	db.Model(Users{}).Select("Name").Where("uid IN (?)", db.Model(Beds{}).Select("uid").Where("room_id = ?", roomId)).Scan(&names)

	for _, name := range names {
		m := map[string]string{"name": name}
		roomMate = append(roomMate, m)
	}

	return
}

type EmptyBedsApi struct {
	BuildingId int `json:"building_id" gorm:"building_id"`
	Gender     int `json:"gender"`
	Cnt        int `json:"cnt"`
}

// GetEmptyBeds : get all empty beds grouped by building's id according to gender
func GetEmptyBeds(gender int) (list []EmptyBedsApi) {
	var db = database.MysqlDb

	// get all valid building's id
	var bIds []int
	db.Model(Rooms{}).Select("building_id").Distinct("building_id").Where("is_valid = ? and gender = ? and building_id IN (?)", 1, gender, db.Model(Buildings{}).Select("building_id").Where("is_valid = ? and is_del = ?", 1, 0)).Scan(&bIds)

	// calculate every building's empty beds
	for _, bId := range bIds {
		// bId : current building's id

		// cnt : all empty beds. gorm required int64
		var cnt int64
		cnt = 0
		var roomIds []int
		db.Model(Rooms{}).Select("id").Where("building_id = ? and gender = ? and is_valid = ? and is_del = ?", bId, gender, 1, 0).Scan(&roomIds)
		// calculate every room's empty beds
		for _, roomId := range roomIds {
			var bedCnt int64
			db.Model(Beds{}).Where("is_valid = ? and is_del = ? and status = ? and room_id = ?", 1, 0, 0, roomId).Count(&bedCnt)
			_ = cache.SetRoomCache(roomId, int(bedCnt))
			cnt += bedCnt
		}

		// set building cache
		_ = cache.SetBuildingCache(bId, gender, int(cnt))

		emptyBeds := EmptyBedsApi{
			BuildingId: bId,
			Gender:     gender,
			Cnt:        int(cnt),
		}
		list = append(list, emptyBeds)
	}
	return
}
