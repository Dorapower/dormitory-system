package model

import "dormitory-system/src/database"

type User struct {
	Uid         int `gorm:"primaryKey;autoIncrement"`
	Name        string
	Gender      int
	Email       string `gorm:"default:NULL"`
	Mobile      string `gorm:"default:NULL"`
	Type        int    `gorm:"default:1"`
	AddedAt     int
	DeletedAt   int    `gorm:"default:NULL"`
	DeniedAt    int    `gorm:"default:NULL"`
	LastLoginAt int    `gorm:"default:NULL"`
	Remarks     string `gorm:"default:NULL"`
	Status      int    `gorm:"default:0"`
}

func GetUserByUid(uid int) (user User) {
	var db = database.MysqlDb
	db.Where("uid = ?", uid).First(&user)
	return
}

// update last_login_at
func (u *User) updateLastLogin() bool {
	var db = database.MysqlDb
	result := db.Model(User{}).Where("uid = ?", u.Uid).Update("last_login_at", u.LastLoginAt)
	if result.Error != nil {
		return false
	}
	return true
}
