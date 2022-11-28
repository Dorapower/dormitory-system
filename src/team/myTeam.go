package team

import (
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
)

func MyTeamHandler(ctx *gin.Context) {
	uid := ctx.Keys["uid"].(int)
	teams := model.GetMyGroup(uid)
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "get my teams success",
		"data":    teams,
	})
}
