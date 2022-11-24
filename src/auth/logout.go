package auth

import (
	"dormitory-system/src/cache"
	"github.com/gin-gonic/gin"
)

func LogoutHandler(ctx *gin.Context) {
	uid := ctx.Keys["uid"].(int)
	if err := cache.DeleteRefreshTokenCache(uid); err != nil {
		ctx.JSON(500, gin.H{
			"error_code": 1,
			"message":    "server error when deleting token",
			"data":       gin.H{},
		})
		return
	}
	ctx.JSON(200, gin.H{
		"error_code": 0,
		"message":    "logout success",
		"data":       gin.H{},
	})
}
