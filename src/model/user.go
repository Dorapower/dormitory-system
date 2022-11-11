package model

import "dormitory-system/src/database"

type User struct {
	Uid         int `gorm:"primaryKey"`
	Name        string
	Gender      int
	Email       string
	Mobile      string
	Type        int
	AddedAt     int
	DeletedAt   int
	DeniedAt    int
	LastLoginAt int
	Remarks     string
	Status      int
}

func getUserByUid(uid int) (user User) {
	var db = database.MysqlDb
	db.Where("uid = ?", uid).First(&user)
	return
}
