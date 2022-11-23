package model

import "dormitory-system/src/database"

type Class struct {
	ID      int `gorm:"primaryKey;autoIncrement"`
	Name    string
	Remarks string
	IsDel   int `gorm:"default:0"`
	Status  int `gorm:"default:0"`
}

func GetClassName(id int) string {
	var db = database.MysqlDb
	var class Class
	db.First(&class, id)
	return class.Name
}
