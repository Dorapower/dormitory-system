package team

import (
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
)

type JoinTeamRequest struct {
	InviteCode string `form:"invite_code" json:"invite_code" binding:"required"`
}

func JoinTeamHandler(ctx *gin.Context) {
	// join teams based on requests
	uid := ctx.Keys["uid"].(int)
	var joinTeamRequest JoinTeamRequest
	err := ctx.ShouldBind(&joinTeamRequest)
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    1,
			"message": "bad request",
			"data":    gin.H{},
		})
		return
	}
	result := model.JoinGroup(uid, joinTeamRequest.InviteCode)
	switch result {
	case 0:
		ctx.JSON(200, gin.H{
			"code":    200,
			"message": "join team success",
			"data":    gin.H{},
		})
	case 1:
		ctx.JSON(500, gin.H{
			"code":    500,
			"message": "invite code wrong",
			"data":    gin.H{},
		})
	case 2:
		ctx.JSON(500, gin.H{
			"code":    500,
			"message": "gender does not match",
			"data":    gin.H{},
		})
	case 3:
		ctx.JSON(500, gin.H{
			"code":    500,
			"message": "target team is full",
			"data":    gin.H{},
		})
	}
}
