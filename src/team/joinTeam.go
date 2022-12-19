package team

import (
	"dormitory-system/src/model"
	"dormitory-system/statuscode"
	"github.com/gin-gonic/gin"
	"net/http"
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    statuscode.StatusInvalidRequest,
			"message": "bad request",
			"data":    gin.H{},
		})
		return
	}
	result := model.JoinGroup(uid, joinTeamRequest.InviteCode)
	switch result {
	case 0:
		ctx.JSON(http.StatusOK, gin.H{
			"code":    statuscode.StatusSuccess,
			"message": "join team success",
			"data":    gin.H{},
		})
	case 1:
		ctx.JSON(http.StatusOK, gin.H{
			"code":    statuscode.StatusJoinTeamFailed,
			"message": "already have a group",
			"data":    gin.H{},
		})
	case 2:
		ctx.JSON(http.StatusOK, gin.H{
			"code":    statuscode.StatusInvalidInviteCode,
			"message": "invite code wrong",
			"data":    gin.H{},
		})
	case 3:
		ctx.JSON(http.StatusOK, gin.H{
			"code":    statuscode.StatusMisMatchedGender,
			"message": "gender does not match",
			"data":    gin.H{},
		})
	case 4:
		ctx.JSON(http.StatusOK, gin.H{
			"code":    statuscode.StatusTeamIsFull,
			"message": "target team is full",
			"data":    gin.H{},
		})
	}
}
