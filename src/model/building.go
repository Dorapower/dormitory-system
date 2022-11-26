package model

import "dormitory-system/src/database"

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
	building_id   int
	building_name string
}

// GetBuildingList : get all building's id and name
func GetBuildingList() (list []BuildingApi) {
	var db = database.MysqlDb
	db.Model(Buildings{}).Order("order_num").Select("id", "Name").Scan(&list)
	return
}

type BuildingInfoApi struct {
	name      string
	describe  string
	image_url string
}

// GetBuildingInfo : get a building's detailed information
func GetBuildingInfo(id int) (list []BuildingInfoApi) {
	var db = database.MysqlDb
	db.Model(Buildings{}).Select("name", "describe", "image_url").Where("id = ?", id).Scan(&list)
	return
}
