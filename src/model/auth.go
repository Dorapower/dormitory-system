package model

import (
	"crypto/md5"
	"dormitory-system/src/database"
	"encoding/hex"
)

type Auth struct {
	Aid       int `gorm:"primaryKey;autoIncrement"`
	Type      int `gorm:"default:0"`
	Username  string
	Password  string
	Salt      string
	Uid       int
	AddedAt   int
	DeletedAt int    `gorm:"default:NULL"`
	Remarks   string `gorm:"default:NULL"`
	Status    int    `gorm:"default:0"`
}

func CheckAuth(username, password string, type_ int) User {
	var db = database.MysqlDb
	var auth Auth
	result := db.Where("username = ? and type = ?", username, type_).Find(&auth)

	// gorm.ErrRecordNotFound
	if result.Error != nil {
		return User{}
	}

	// password wrong
	if auth.Password != getPwd(password, auth.Salt) {
		return User{}
	}

	var user User
	user = GetUserByUid(auth.Uid)
	//	user.updateLastLogin()
	return user
}

// get encrypted password
func getPwd(str, salt string) string {
	b := []byte(str)
	s := []byte(salt)
	h := md5.New()
	h.Write(s)
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}
