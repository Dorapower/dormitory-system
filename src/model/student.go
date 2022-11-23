package model

import "dormitory-system/src/database"

type StudentInfo struct {
	ID               int `gorm:"primaryKey;autoIncrement"`
	Uid              int
	StudentId        string `gorm:"column:studentid"`
	VerificationCode string `gorm:"default:0"`
	ClassId          int
	Remarks          string
	Status           int `gorm:"default:0"`
}

func GetStudentByUid(uid int) (stu StudentInfo) {
	var db = database.MysqlDb
	db.Where("uid = ?", uid).First(&stu)
	return
}
