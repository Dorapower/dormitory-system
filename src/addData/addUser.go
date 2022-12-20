package addData

import (
	"dormitory-system/src/database"
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"time"
)

func AddUser(ctx *gin.Context) {
	var user User
	err := ctx.MustBindWith(&user, binding.JSON)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}
	db := database.MysqlDb
	var userInfo = model.Users{
		Name:          user.Name,
		Gender:        user.Gender,
		Email:         user.Email,
		Tel:           user.Tel,
		Type:          user.Type,
		AddTime:       int(time.Now().Unix()),
		IsDel:         user.IsDel,
		IsDeny:        user.IsDeny,
		LastLoginTime: user.LastLoginTime,
		Remarks:       user.Remarks,
		Status:        user.Status,
	}
	db.Create(&userInfo)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "增加成功",
	})
	return
}
