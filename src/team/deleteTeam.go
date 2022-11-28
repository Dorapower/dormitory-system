package team

import (
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
)

type DeleteTeamRequest struct {
	TeamId int `form:"team_id" json:"team_id" binding:"required"`
}

func DeleteTeamHandler(ctx *gin.Context) {
	// delete teams based on requests
	// uid := ctx.Keys["uid"].(int)
	var deleteTeamRequest DeleteTeamRequest
	err := ctx.ShouldBind(&deleteTeamRequest)
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    1,
			"message": "bad request",
			"data":    gin.H{},
		})
		return
	}
	success := model.DelGroup(deleteTeamRequest.TeamId)
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
