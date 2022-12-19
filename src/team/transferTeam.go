package team

import (
	"dormitory-system/src/model"
	"dormitory-system/statuscode"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransferTeamRequest struct {
	StudentId string `form:"student_id" json:"student_id" binding:"required"`
}

func TransferTeamHandler(ctx *gin.Context) {
	uid := ctx.Keys["uid"].(int)
	var transferTeamRequest TransferTeamRequest
	err := ctx.ShouldBind(&transferTeamRequest)
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    1,
			"message": "bad request",
			"data":    gin.H{},
		})
		return
	}
	success := model.TransferGroup(uid, transferTeamRequest.StudentId)
	if success {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    statuscode.StatusSuccess,
			"message": "transfer team success",
			"data":    gin.H{},
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    statuscode.StatusServerError,
			"message": "transfer team failed",
			"data":    gin.H{},
		})
	}
}
