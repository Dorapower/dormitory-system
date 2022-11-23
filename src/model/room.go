package model

import "dormitory-system/src/database"

type Room struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	BuildingId int
	Name       string
	Gender     int
	OrderNum   int
	IsValid    int `gorm:"default:1"`
	Remarks    string
	Describe   string
	ImageUrl   string
	IsDel      int `gorm:"default:0"`
	Status     int `gorm:"default:0"`
}

func GetRoomInfoById(id int) (room Room) {
	var db = database.MysqlDb
	db.First(&room, id)
	return
}
