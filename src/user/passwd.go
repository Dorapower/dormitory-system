package user

import (
	"dormitory-system/src/model"
	"dormitory-system/statuscode"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error_code": statuscode.StatusInvalidRequest,
			"message":    "bad request",
		})
	}
	uid := ctx.Keys["uid"].(int)
	if model.ChangePwd(uid, passwordRequest.OldPassword, passwordRequest.NewPassword) == false {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    statuscode.StatusWrongPassword,
			"message": "old password is wrong",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    statuscode.StatusSuccess,
		"message": "change password success",
	})
}
