package middleware

import (
	"dormitory-system/src/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    1,
				"message": "missing token",
			})
			ctx.Abort()
			return
		}
		uid, err := auth.ParseToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    2,
				"message": "invalid token:" + err.Error(),
			})
			ctx.Abort()
			return
		}
		ctx.Set("uid", uid)
		ctx.Next()
	}
}
