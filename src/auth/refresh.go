package auth

import (
	"dormitory-system/src/cache"
	"dormitory-system/src/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
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
		return os.Getenv("API_SECRET"), nil
	}, jwt.WithJSONNumber())
	if err != nil {
		ctx.JSON(400, gin.H{
			"error_code": 2,
			"message":    "invalid token:" + err.Error(),
			"data":       gin.H{},
		})
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if claims["exp"] == nil || claims["uid"] == nil {
			ctx.JSON(400, gin.H{
				"error_code": 3,
				"message":    "invalid refresh token",
				"data":       gin.H{},
			})
			return
		}
		expireAt, err := claims["exp"].(json.Number).Int64()
		if err != nil {
			log.Println(err)
			ctx.JSON(500, gin.H{
				"error_code": 4,
				"message":    "server error when parsing token",
				"data":       gin.H{},
			})
			return
		}
		rawUid, err := claims["uid"].(json.Number).Int64()
		uid := int(rawUid)
		if err != nil {
			log.Println(err)
			ctx.JSON(500, gin.H{
				"error_code": 4,
				"message":    "server error when parsing token",
				"data":       gin.H{},
			})
			return
		}
		if time.Unix(expireAt, 0).Before(time.Now()) {
			ctx.JSON(400, gin.H{
				"error_code": 5,
				"message":    "refresh token expired",
				"data":       gin.H{},
			})
			return
		}
		if cache.GetRefreshTokenCache(uid) != refreshRequest.RefreshToken {
			log.Println(uid)
			log.Println(cache.GetRefreshTokenCache(uid))
			ctx.JSON(400, gin.H{
				"error_code": 6,
				"message":    "outdated refresh token",
				"data":       gin.H{},
			})
			return
		}
		user := model.GetUserByUid(uid)
		tokenString, refreshTokenString, err := generateTokenPair(&user)
		if err != nil {
			ctx.JSON(400, gin.H{
				"error_code": 7,
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
		err = cache.SetRefreshTokenCache(refreshTokenString, uid)
		if err != nil {
			log.Println(err)
		}
	} else {
		ctx.JSON(400, gin.H{
			"error_code": 8,
			"message":    "invalid refresh token",
			"data":       gin.H{},
		})
		return
	}

}
