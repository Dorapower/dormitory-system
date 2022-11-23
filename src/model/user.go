package model

import "dormitory-system/src/database"

type User struct {
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

func GetUserByUid(uid int) (user User) {
	var db = database.MysqlDb
	db.Where("uid = ?", uid).First(&user)
	return
}

// update last_login_at
func (u *User) updateLastLogin() bool {
	var db = database.MysqlDb
	result := db.Model(User{}).Where("uid = ?", u.Uid).Update("last_login_at", u.LastLoginTime)
	if result.Error != nil {
		return false
	}
	return true
}
