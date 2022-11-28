package user

import (
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
)

func MyInfoHandler(ctx *gin.Context) {
	uid := ctx.Keys["uid"].(int)
	user := model.GetUserByUid(uid)
	if user == (model.Users{}) {
		ctx.JSON(500, gin.H{
			"error_code": 1,
			"message":    "server error when getting user info",
			"data":       gin.H{},
		})
		return
	}
	ctx.JSON(200, gin.H{
		"error_code": 0,
		"message":    "get user info success",
		"data":       user,
	})
}
