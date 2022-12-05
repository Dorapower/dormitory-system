package user

import (
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
)

func MyInfoHandler(ctx *gin.Context) {
	uid := ctx.Keys["uid"].(int)
	user, err := model.GetUserInfoByUid(uid)
	if err != nil {
		ctx.JSON(500, gin.H{
			"code":    1,
			"message": "server error when getting user info",
			"data":    gin.H{},
		})
	}
	if user, ok := user.(model.UserApi); ok {
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "get user info success",
			"data":    user,
		})
	} else {
		ctx.JSON(500, gin.H{
			"code":    2,
			"message": "server error when parsing user info",
			"data":    gin.H{},
		})
	}
}
