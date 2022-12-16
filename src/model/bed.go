package model

import (
	"database/sql"
	"dormitory-system/src/database"
	"log"
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

// GetMyRoomByUid : get my room's Name and roommates' Name
func GetMyRoomByUid(uid int) (roomId int, roomName string, names []string) {
	var db = database.MysqlDb

	// get room's id
	db.Model(Beds{}).Select("room_id").Where("uid = ? and is_valid = ? and is_del = ? and status = ?", uid, 1, 0, 1).Scan(&roomId)

	// get room's Name
	db.Model(Rooms{}).Select("Name").Where("id = ?", roomId).Scan(&roomName)

	// get roommate's Name
	db.Model(Users{}).Select("Name").Where("uid IN (?)", db.Model(Beds{}).Select("uid").Where("room_id = ?", roomId)).Scan(&names)

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
	rows, _ := db.Model(Rooms{}).Select("building_id").Distinct("building_id").Where("is_valid = ? and gender = ? and building_id IN (?)", 1, gender, db.Model(Buildings{}).Select("building_id").Where("is_valid = ?", 1)).Rows()
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)

	// calculate every building's empty beds
	for rows.Next() {
		// bId : current building's id
		var bId int
		err := db.ScanRows(rows, &bId)
		if err != nil {
			return nil
		}

		// cnt : all empty beds. gorm required int64
		var cnt int64
		db.Model(Beds{}).Where("is_valid = ? and is_del = ? and status = ? and room_id IN (?)", 1, 0, 0, db.Model(Rooms{}).Select("id").Where("gender = ? and building_id = ?", gender, bId)).Count(&cnt)

		emptyBeds := EmptyBedsApi{
			BuildingId: bId,
			Gender:     gender,
			Cnt:        int(cnt),
		}
		list = append(list, emptyBeds)
	}
	return
}
