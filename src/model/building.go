package model

import (
	"dormitory-system/src/database"
)

type Buildings struct {
	ID       int `gorm:"primaryKey;autoIncrement"`
	Name     string
	OrderNum int
	IsValid  int `gorm:"default:1"`
	Remarks  string
	Describe string
	ImageUrl string
	IsDel    int `gorm:"default:0"`
	Status   int `gorm:"default:0"`
}

type BuildingApi struct {
	BuildingId   int    `json:"building_id" gorm:"column:id"`
	BuildingName string `json:"building_name" gorm:"column:name"`
}

// GetBuildingList : get all building's id and Name
func GetBuildingList() (list []BuildingApi) {
	var db = database.MysqlDb
	db.Model(Buildings{}).Order("order_num").Select("id", "name").Scan(&list)
	return
}

type BuildingInfoApi struct {
	Name     string `json:"name"`
	Describe string `json:"describe"`
	ImageUrl string `json:"image_url"`
}

// GetBuildingInfo : get a building's detailed information
func GetBuildingInfo(id int) (bIA BuildingInfoApi) {
	var db = database.MysqlDb
	db.Model(Buildings{}).Select("name", "describe", "image_url").Where("id = ?", id).Scan(&bIA)
	return
}
