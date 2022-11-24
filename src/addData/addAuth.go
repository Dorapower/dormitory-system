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

func AddAuth(ctx *gin.Context) {
	var auth Auth
	err := ctx.MustBindWith(&auth, binding.JSON)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	db := database.MysqlDb
	var authInfo = model.Auth{
		Username: auth.Username,
		Password: auth.Password,
		Uid:      auth.Uid,
		AddTime:  int(time.Now().Unix()),
		Remarks:  auth.Remarks,
		Status:   auth.Status,
	}
	db.Create(&authInfo)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "增加成功",
	})
	return
}
