package addData

import (
	"crypto/md5"
	"dormitory-system/src/database"
	"dormitory-system/src/model"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"math/rand"
	"net/http"
	"time"
)

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

func AddAuth(ctx *gin.Context) {
	var auth Auth
	ctx.MustBindWith(&auth, binding.JSON)
	db := database.MysqlDb
	salt := getSalt()
	var authInfo = model.Auth{
		Type:      auth.Type,
		Username:  auth.Username,
		Salt:      salt,
		Password:  getPwd(auth.Password, salt),
		Uid:       auth.Uid,
		AddedAt:   int(time.Now().Unix()),
		DeletedAt: auth.DeletedAt,
		Remarks:   auth.Remarks,
		Status:    auth.Status,
	}
	db.Create(&authInfo)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "增加成功",
	})
	return
}
