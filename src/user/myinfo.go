package user

import (
	"dormitory-system/src/model"
	"dormitory-system/statuscode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MyInfoHandler(ctx *gin.Context) {
	uid := ctx.Keys["uid"].(int)
	user, err := model.GetUserInfoByUid(uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    statuscode.StatusServerError,
			"message": "server error when getting user info",
			"data":    gin.H{},
		})
	}
	if user, ok := user.(model.UserApi); ok {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    statuscode.StatusSuccess,
			"message": "get user info success",
			"data":    user,
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    statuscode.StatusServerError,
			"message": "server error when parsing user info",
			"data":    gin.H{},
		})
	}
}
