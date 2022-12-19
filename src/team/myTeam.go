package team

import (
	"dormitory-system/src/model"
	"dormitory-system/statuscode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MyTeamHandler(ctx *gin.Context) {
	uid := ctx.Keys["uid"].(int)
	teams := model.GetMyGroup(uid)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    statuscode.StatusSuccess,
		"message": "get my teams success",
		"data":    teams,
	})
}
