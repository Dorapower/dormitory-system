package team

import (
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
)

type CreateTeamRequest struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Describe string `form:"describe" json:"describe" binding:"required"`
}

func CreateTeamHandler(ctx *gin.Context) {
	// create teams based on requests
	var createTeamRequest CreateTeamRequest
	err := ctx.ShouldBind(&createTeamRequest)
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    1,
			"message": "bad request",
			"data":    gin.H{},
		})
		return
	}
	uid := ctx.Keys["uid"].(int)
	created := model.CreatGroup(uid, createTeamRequest.Name, createTeamRequest.Describe)
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "create team success",
		"data": gin.H{
			"team_id":     created.TeamId,
			"invite_code": created.InviteCode,
		},
	})
}
