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
	uid               int
	studengid         string
	name              string
	gender            int
	email             string
	tel               string
	last_login_time   string
	verification_code string
	class_name        string
}

func GetUserInfoByUid(uid int) (userApi UserApi) {
	var db = database.MysqlDb
	var user Users
	db.Where("uid = ?", uid).First(&user)
	var stu StudentInfo
	db.Where("uid = ?", uid).First(&stu)
	className := GetClassName(stu.ClassId)

	lastLoginTime := time.Unix(int64(user.LastLoginTime), 0).Format("2006-01-02 15:04:05")

	userApi.uid = uid
	userApi.studengid = stu.StudentId
	userApi.name = user.Name
	userApi.gender = user.Gender
	userApi.email = user.Email
	userApi.tel = user.Tel
	userApi.last_login_time = lastLoginTime
	userApi.verification_code = stu.VerificationCode
	userApi.class_name = className

	return
}

// update last_login_at
func (u *Users) updateLastLogin() bool {
	var db = database.MysqlDb
	result := db.Model(Users{}).Where("uid = ?", u.Uid).Update("last_login_at", u.LastLoginTime)
	if result.Error != nil {
		return false
	}
	return true
}
