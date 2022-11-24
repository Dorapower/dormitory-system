package model

import (
	"dormitory-system/src/database"
)

type Beds struct {
	ID       int `gorm:"primaryKey;autoIncrement"`
	Uid      int `gorm:"default:0"`
	RoomId   int
	Name     int
	OrderNum int
	IsValid  int `gorm:"default:1"`
	Remarks  string
	IsDel    int `gorm:"default:0"`
	Status   int `gorm:"default:0"`
}

// GetMyRoomByUid : get my room's name and roommates' name
func GetMyRoomByUid(uid int) (roomName string, names []string) {
	var db = database.MysqlDb

	var roomId int
	db.Model(Beds{}).Select("room_id").Where("uid = ?", uid).Scan(&roomId)

	// get room's name
	db.Model(Rooms{}).Select("name").Where("id = ?", roomId).Scan(&roomName)

	// get roommate's name
	db.Model(User{}).Select("name").Where("uid IN (?)", db.Model(Beds{}).Select("uid").Where("room_id = ?", roomId)).Scan(&names)

	return
}

type EmptyBedsApi struct {
	building_id int
	gender      int
	cnt         int
}

func GetEmptyBeds(gender int) (list []EmptyBedsApi) {
	var db = database.MysqlDb

	// get all valid building's id
	rows, _ := db.Model(Rooms{}).Select("building_id").Distinct("building_id").Where("gender = ? and building_id IN (?)", gender, db.Model(Buildings{}).Select("building_id").Where("is_valid = ?", 1)).Rows()
	defer rows.Close()

	// calculate every building's empty beds
	for rows.Next() {
		// bId : current building's id
		var bId int
		db.ScanRows(rows, &bId)

		// cnt : all empty beds. gorm required int64
		var cnt int64
		db.Model(Beds{}).Where("is_valid = ? and room_id IN (?)", 1, db.Model(Rooms{}).Select("id").Where("gender = ? and building_id = ?", gender, bId)).Count(&cnt)

		emptyBeds := EmptyBedsApi{
			building_id: bId,
			gender:      gender,
			cnt:         int(cnt),
		}
		list = append(list, emptyBeds)
	}
	return
}
