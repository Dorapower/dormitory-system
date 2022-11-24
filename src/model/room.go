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

type RoomInfoApi struct {
	name        string
	gender      int
	describe    string
	image_url   string
	building_id int
}

// GetRoomInfoById : get room's detailed information
func GetRoomInfoById(id int) (rApi RoomInfoApi) {
	var db = database.MysqlDb
	db.Model(Room{}).Select("name", "gender", "describe", "image_url", "building_id").Where("id = ?", id).Scan(&rApi)
	return
}
