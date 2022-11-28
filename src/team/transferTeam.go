package team

import (
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
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
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "transfer team success",
			"data":    gin.H{},
		})
	} else {
		ctx.JSON(500, gin.H{
			"code":    500,
			"message": "transfer team failed",
			"data":    gin.H{},
		})
	}
}
