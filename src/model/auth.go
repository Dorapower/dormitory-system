package model

import (
	"dormitory-system/src/database"
)

type Auth struct {
	id int `gorm:"primaryKey;autoIncrement"`
	//Type      int `gorm:"default:0"`
	Username string
	Password string
	//Salt      string
	Uid       int
	AddedTime int `gorm:"default:0"`
	Remarks   string
	IsDel     int `gorm:"default:0"`
	Status    int `gorm:"default:0"`
}

func CheckAuth(username, password string) User { // delete type param
	var db = database.MysqlDb
	var auth Auth
	result := db.Where("username = ? and password = ?", username, password).Find(&auth)

	// gorm.ErrRecordNotFound
	if result.Error != nil {
		return User{}
	}

	// password wrong
	//if auth.Password != getPwd(password, auth.Salt) {
	//	return User{}
	//}

	var user User
	user = GetUserByUid(auth.Uid)
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
	db.Model(&auth).Update("password", newPwd)
	return true
}
