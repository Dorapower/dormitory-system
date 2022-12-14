package team

import (
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
)

func DeleteTeamHandler(ctx *gin.Context) {
	uid := ctx.Keys["uid"].(int)
	success := model.DelGroup(uid)
	if success {
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "delete team success",
			"data":    gin.H{},
		})
	} else {
		ctx.JSON(500, gin.H{
			"code":    500,
			"message": "delete team failed",
			"data":    gin.H{},
		})
	}
}
