package model

import "dormitory-system/src/database"

type Rooms struct {
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
	Name       string `json:"name" gorm:"column:name"`
	Gender     int    `json:"gender" gorm:"column:gender"`
	Describe   string `json:"describe" gorm:"colum:describe"`
	ImageUrl   string `json:"image_url" gorm:"image_url"`
	BuildingId int    `json:"building_id" gorm:"building_id"`
}

// GetRoomInfoById : get room's detailed information
func GetRoomInfoById(id int) (rApi RoomInfoApi) {
	var db = database.MysqlDb
	db.Model(Rooms{}).Select("Name", "gender", "describe", "image_url", "building_id").Where("id = ?", id).Scan(&rApi)
	return
}
