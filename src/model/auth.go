package model

import (
	"dormitory-system/src/database"
	"time"
)

type Auth struct {
	id int `gorm:"primaryKey;autoIncrement"`
	//Type      int `gorm:"default:0"`
	Username string
	Password string
	//Salt      string
	Uid     int
	AddTime int `gorm:"default:0"`
	Remarks string
	IsDel   int `gorm:"default:0"`
	Status  int `gorm:"default:0"`
}

func CheckAuth(username, password string) Users { // delete type param
	var db = database.MysqlDb
	var auth Auth
	result := db.Where("username = ? and password = ?", username, password).First(&auth)

	// gorm.ErrRecordNotFound

	if result.Error != nil {
		return Users{}
	}

	// password wrong
	//if auth.Password != getPwd(password, auth.Salt) {
	//	return User{}
	//}

	var user Users
	user = GetUserByUid(auth.Uid)
	user.LastLoginTime = int(time.Now().Unix())
	user.updateLastLogin()

	return user
}

// get encrypted password
//func getPwd(str, salt string) string {
//	b := []byte(str)
//	s := []byte(salt)
//	h := md5.New()
//	h.Write(s)
//	h.Write(b)
//	return hex.EncodeToString(h.Sum(nil))
//}

func ChangePwd(uid int, oldPwd, newPwd string) bool {
	var db = database.MysqlDb
	var auth Auth
	db.Where("uid = ?", uid).First(&auth)
	if auth.Password != oldPwd {
		return false
	}
	db.Model(&auth).Where("uid = ?", uid).Update("password", newPwd)
	return true
}
