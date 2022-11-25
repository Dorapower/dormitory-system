package auth

import (
	"dormitory-system/src/cache"
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func LoginHandler(ctx *gin.Context) {
	var login Login
	var user model.Users
	if err := ctx.MustBindWith(&login, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error_code": 1,
			"message":    "missing username or password",
			"data":       gin.H{},
		})
		return
	}
	if user = checkLogin(login); user == (model.Users{}) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error_code": 2,
			"message":    "wrong username or password",
			"data":       gin.H{},
		})
		return
	}
	token, refreshToken, err := generateTokenPair(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error_code": 3,
			"message":    "server error when generating tokens",
			"data":       gin.H{},
		})
		return
	}
	err = cache.SetRefreshTokenCache(refreshToken, user.Uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error_code": 4,
			"message":    "server error when caching refresh token",
			"data":       gin.H{},
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error_code": 0,
		"message":    "login success",
		"data": gin.H{
			"token":         token,
			"refresh_token": refreshToken,
			"expires_in":    TokenDuration.Seconds(),
		},
	})
	return
}
func checkLogin(login Login) model.Users {
	user := model.CheckAuth(login.Username, login.Password) // delete type
	return user
}
