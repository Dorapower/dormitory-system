package model

import "dormitory-system/src/database"

type Sys struct {
	ID       int `gorm:"primaryKey;autoIncrement"`
	KeyName  string
	KeyValue string
	IsDel    int `gorm:"default:0"`
	Remarks  string
}

func GetSystemConfigByKey(key string) (sys Sys) {
	var db = database.MysqlDb
	db.Where("key_name = ?", key).First(&sys)
	return
}
