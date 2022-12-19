package auth

import (
	"dormitory-system/src/cache"
	"dormitory-system/src/model"
	"dormitory-system/statuscode"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"os"
	"time"
)

func RefreshHandler(ctx *gin.Context) {
	var refreshRequest RefreshRequest
	err := ctx.BindJSON(&refreshRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error_code": statuscode.StatusNoToken,
			"message":    "missing refresh token",
			"data":       gin.H{},
		})
		return
	}
	token, err := jwt.Parse(refreshRequest.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	}, jwt.WithJSONNumber())
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error_code": statuscode.StatusInvalidToken,
			"message":    "invalid token:" + err.Error(),
		})
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if claims["exp"] == nil || claims["uid"] == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"error_code": statuscode.StatusInvalidToken,
				"message":    "invalid refresh token",
			})
			return
		}
		expireAt, err := claims["exp"].(json.Number).Int64()
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error_code": statuscode.StatusServerError,
				"message":    "server error when parsing token",
			})
			return
		}
		rawUid, err := claims["uid"].(json.Number).Int64()
		uid := int(rawUid)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error_code": statuscode.StatusServerError,
				"message":    "server error when parsing token",
			})
			return
		}
		if time.Unix(expireAt, 0).Before(time.Now()) {
			ctx.JSON(http.StatusOK, gin.H{
				"error_code": statuscode.StatusExpiredToken,
				"message":    "refresh token expired",
			})
			return
		}
		if cache.GetRefreshTokenCache(uid) != refreshRequest.RefreshToken {
			log.Println(uid)
			log.Println(cache.GetRefreshTokenCache(uid))
			ctx.JSON(http.StatusOK, gin.H{
				"error_code": statuscode.StatusInvalidToken,
				"message":    "outdated refresh token",
			})
			return
		}
		user := model.GetUserByUid(uid)
		tokenString, refreshTokenString, err := generateTokenPair(&user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error_code": statuscode.StatusServerError,
				"message":    "failed to generate token pair",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"error_code": statuscode.StatusSuccess,
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error_code": statuscode.StatusInvalidToken,
			"message":    "invalid refresh token",
		})
		return
	}

}
