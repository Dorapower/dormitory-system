package auth

import (
	"dormitory-system/src/cache"
	"dormitory-system/src/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

func RefreshHandler(ctx *gin.Context) {
	var refreshRequest RefreshRequest
	err := ctx.BindJSON(&refreshRequest)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error_code": 1,
			"message":    "missing refresh token",
			"data":       gin.H{},
		})
		return
	}
	token, err := jwt.Parse(refreshRequest.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if time.Unix(claims["exp"].(int64), 0).Before(time.Now()) {
			ctx.JSON(400, gin.H{
				"error_code": 2,
				"message":    "refresh token expired",
				"data":       gin.H{},
			})
			return
		}
		if claims["uid"] == nil {
			ctx.JSON(400, gin.H{
				"error_code": 3,
				"message":    "invalid refresh token",
				"data":       gin.H{},
			})
			return
		}
		if cache.GetRefreshTokenCache(claims["uid"].(int)) != refreshRequest.RefreshToken {
			ctx.JSON(400, gin.H{
				"error_code": 4,
				"message":    "invalid refresh token",
				"data":       gin.H{},
			})
			return
		}
		user := model.GetUserByUid(claims["uid"].(int))
		tokenString, refreshTokenString, err := generateTokenPair(&user)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error_code": 5,
				"message":    "failed to generate token pair",
				"data":       gin.H{},
			})
			return
		}
		ctx.JSON(200, gin.H{
			"error_code": 0,
			"message":    "refresh token pair successfully",
			"data": gin.H{
				"token":         tokenString,
				"refresh_token": refreshTokenString,
			},
		})
		err = cache.SetRefreshTokenCache(refreshTokenString, claims["uid"].(int))
		if err != nil {
			log.Println(err)
		}
	} else {
		ctx.JSON(400, gin.H{
			"error_code": 6,
			"message":    "invalid refresh token",
			"data":       gin.H{},
		})
		return
	}

}
