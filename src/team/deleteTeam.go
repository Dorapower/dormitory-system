package team

import (
	"dormitory-system/src/model"
	"dormitory-system/statuscode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteTeamHandler(ctx *gin.Context) {
	uid := ctx.Keys["uid"].(int)
	success := model.DelGroup(uid)
	if success {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    statuscode.StatusSuccess,
			"message": "delete team success",
			"data":    gin.H{},
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    statuscode.StatusServerError,
			"message": "delete team failed",
			"data":    gin.H{},
		})
	}
}
