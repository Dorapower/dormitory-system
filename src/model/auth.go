package model

import (
	"crypto/md5"
	"dormitory-system/src/database"
	"encoding/hex"
	"math/rand"
	"time"
)

type Auth struct {
	Aid       int `gorm:"primaryKey"`
	Type      int
	Username  string
	Password  string
	salt      string
	Uid       int
	AddedAt   int
	DeletedAt int
	Remarks   string
	Status    int
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
	if auth.Password != getPwd(password, auth.salt) {
		return User{}
	}

	var user User
	user = getUserByUid(auth.Uid)
	//	user.updateLastLogin()
	return user
}

// generate a salt
func getSalt() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
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
