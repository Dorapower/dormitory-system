package auth

import (
	"dormitory-system/src/cache"
	"dormitory-system/statuscode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LogoutHandler(ctx *gin.Context) {
	uid := ctx.Keys["uid"].(int)
	if err := cache.DeleteRefreshTokenCache(uid); err != nil {
		ctx.JSON(500, gin.H{
			"code":    1,
			"message": "server error when deleting token",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    statuscode.StatusSuccess,
		"message": "logout success",
	})
}
