package addData

import (
	"dormitory-system/src/database"
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"time"
)

func AddUser(ctx *gin.Context) {
	var user User
	ctx.MustBindWith(&user, binding.JSON)
	db := database.MysqlDb
	var userInfo = model.User{
		Name:        user.Name,
		Gender:      user.Gender,
		Email:       user.Email,
		Mobile:      user.Mobile,
		Type:        user.Type,
		AddedAt:     int(time.Now().Unix()),
		DeletedAt:   user.DeletedAt,
		DeniedAt:    user.DeniedAt,
		LastLoginAt: user.LastLoginAt,
		Remarks:     user.Remarks,
		Status:      user.Status,
	}
	db.Create(&userInfo)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "增加成功",
	})
	return
}
