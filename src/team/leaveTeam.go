package team

import (
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
)

type LeaveTeamRequest struct {
	TeamId int `form:"team_id" json:"team_id" binding:"required"`
}

func LeaveTeamHandler(ctx *gin.Context) {
	uid := ctx.Keys["uid"].(int)
	var leaveTeamRequest LeaveTeamRequest
	err := ctx.ShouldBind(&leaveTeamRequest)
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    1,
			"message": "bad request",
			"data":    gin.H{},
		})
		return
	}
	success := model.QuitGroup(uid, leaveTeamRequest.TeamId)
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
