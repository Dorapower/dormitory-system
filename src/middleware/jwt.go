package middleware

import (
	"dormitory-system/src/auth"
	"dormitory-system/statuscode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    statuscode.StatusNoToken,
				"message": "missing token",
			})
			ctx.Abort()
			return
		}
		uid, err := auth.ParseToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    statuscode.StatusInvalidToken,
				"message": "invalid token:" + err.Error(),
			})
			ctx.Abort()
			return
		}
		ctx.Set("uid", uid)
		ctx.Next()
	}
}
