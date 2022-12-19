package team

import (
	"dormitory-system/src/model"
	"dormitory-system/statuscode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LeaveTeamHandler(ctx *gin.Context) {
	uid := ctx.Keys["uid"].(int)
	success := model.QuitGroup(uid)
	if success {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    statuscode.StatusSuccess,
			"message": "leave team success",
			"data":    gin.H{},
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    statuscode.StatusServerError,
			"message": "leave team failed",
			"data":    gin.H{},
		})
	}

}
