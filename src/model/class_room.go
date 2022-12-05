package model

import "dormitory-system/src/database"

type ClassRoom struct {
	ID      int `gorm:"primaryKey;autoIncrement"`
	ClassId int
	RoomId  int
	Remarks string
	IsDel   int `gorm:"default:0"`
	Status  int `gorm:"default:0"`
}

func GetClassRooms(uid int) (roomIds []int) {
	var db = database.MysqlDb
	var classId int
	var gender int
	db.Model(StudentInfo{}).Select("class_id").Where("uid = ?", uid).Scan(&classId)
	db.Model(Users{}).Select("gender").Where("uid = ?", uid).Scan(&gender)
	var tempRoomIds []int
	db.Model(ClassRoom{}).Select("room_id").Where("class_id = ? and is_del = ?", classId, 0).Scan(&tempRoomIds)
	db.Model(Rooms{}).Select("id").Where("id IN (?) and gender = ?", tempRoomIds, gender).Scan(&roomIds)
	return
}
