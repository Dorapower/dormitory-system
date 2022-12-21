package model

import (
	"dormitory-system/src/database"
	"time"
)

type Users struct {
	Uid           int `gorm:"primaryKey;autoIncrement"`
	Name          string
	Gender        int
	Email         string
	Tel           string
	Type          int `gorm:"default:1"`
	AddTime       int `gorm:"default:0"`
	IsDeny        int `gorm:"default:0"`
	LastLoginTime int `gorm:"default:0"`
	Remarks       string
	IsDel         int `gorm:"default:0"`
	Status        int `gorm:"default:0"`
}

func GetUserByUid(uid int) (user Users) {
	var db = database.MysqlDb
	db.Where("uid = ?", uid).First(&user)
	return
}

type UserApi struct {
	Uid              int    `json:"uid"`
	Student          string `json:"studentid"`
	Name             string `json:"name"`
	Gender           int    `json:"gender"`
	Email            string `json:"email"`
	Tel              string `json:"tel"`
	LastLoginTime    string `json:"last_login_time"`
	VerificationCode string `json:"verification_code"`
	ClassName        string `json:"class_name"`
}

func GetUserInfoByUid(uid int) (DbQueryResult, error) {
	var db = database.MysqlDb
	var user Users

	if err := db.Where("uid = ?", uid).First(&user).Error; err != nil {
		return nil, err
	}
	var stu StudentInfo
	if err := db.Where("uid = ?", uid).First(&stu).Error; err != nil {
		return nil, err
	}
	className := GetClassName(stu.ClassId)

	lastLoginTime := time.Unix(int64(user.LastLoginTime), 0).Format("2006-01-02 15:04:05")

	var userApi UserApi

	userApi.Uid = uid
	userApi.Student = stu.StudentId
	userApi.Name = user.Name
	userApi.Gender = user.Gender
	userApi.Email = user.Email
	userApi.Tel = user.Tel
	userApi.LastLoginTime = lastLoginTime
	userApi.VerificationCode = stu.VerificationCode
	userApi.ClassName = className

	return userApi, nil
}

// update last_login_at
func (u *Users) updateLastLogin() bool {
	var db = database.MysqlDb
	result := db.Model(Users{}).Where("uid = ?", u.Uid).Update("last_login_time", u.LastLoginTime)
	if result.Error != nil {
		return false
	}
	return true
}
