package model

import "dormitory-system/src/database"

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
	db.Model(Room{}).Select("name").Where("id = ?", roomId).Scan(&roomName)

	// get roommate's name
	db.Model(User{}).Select("name").Where("uid IN (?)", db.Model(Beds{}).Select("uid").Where("room_id = ?", roomId)).Scan(&names)

	return
}
