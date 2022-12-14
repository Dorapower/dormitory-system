package team

import (
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
)

func LeaveTeamHandler(ctx *gin.Context) {
	uid := ctx.Keys["uid"].(int)
	success := model.QuitGroup(uid)
	if success {
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "leave team success",
			"data":    gin.H{},
		})
	} else {
		ctx.JSON(500, gin.H{
			"code":    500,
			"message": "leave team failed",
			"data":    gin.H{},
		})
	}

}
