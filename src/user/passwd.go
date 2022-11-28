package user

import (
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type PasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

func PasswdHandler(ctx *gin.Context) {
	// users change password
	var passwordRequest PasswordRequest
	err := ctx.MustBindWith(&passwordRequest, binding.JSON)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error_code": 1,
			"message":    "bad request",
			"data":       gin.H{},
		})
	}
	uid := ctx.Keys["uid"].(int)
	if model.ChangePwd(uid, passwordRequest.OldPassword, passwordRequest.NewPassword) == false {
		ctx.JSON(500, gin.H{
			"error_code": 1,
			"message":    "old password is wrong",
			"data":       gin.H{},
		})
		return
	}
	ctx.JSON(200, gin.H{
		"error_code": 0,
		"message":    "change password success",
		"data":       gin.H{},
	})
}
