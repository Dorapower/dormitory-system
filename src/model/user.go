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
