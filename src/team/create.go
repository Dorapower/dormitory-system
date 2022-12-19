package team

import (
	"dormitory-system/src/model"
	"dormitory-system/statuscode"
	"github.com/gin-gonic/gin"
	"net/http"
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    statuscode.StatusInvalidRequest,
			"message": "bad request",
			"data":    gin.H{},
		})
		return
	}
	uid := ctx.Keys["uid"].(int)
	created, ok := model.CreatGroup(uid, createTeamRequest.Name, createTeamRequest.Describe)
	if ok {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    statuscode.StatusSuccess,
			"message": "create team success",
			"data": gin.H{
				"team_id":     created.TeamId,
				"invite_code": created.InviteCode,
			},
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    statuscode.StatusCreateTeamFailed,
			"message": "already have a group",
			"data":    gin.H{},
		})
	}
}
