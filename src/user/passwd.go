package user

import (
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type PasswordRequest struct {
	OldPassword string `json:"oldPasswd" binding:"required"`
	NewPassword string `json:"newPasswd" binding:"required"`
}

func PasswdHandler(ctx *gin.Context) {
	// users change password
	var passwordRequest PasswordRequest
	err := ctx.MustBindWith(&passwordRequest, binding.JSON)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error_code": 1,
			"message":    "bad request",
		})
	}
	uid := ctx.Keys["uid"].(int)
	if model.ChangePwd(uid, passwordRequest.OldPassword, passwordRequest.NewPassword) == false {
		ctx.JSON(500, gin.H{
			"code":    2,
			"message": "old password is wrong",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "change password success",
	})
}
