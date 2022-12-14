package auth

import (
	"dormitory-system/src/cache"
	"dormitory-system/src/model"
	"dormitory-system/statuscode"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func LoginHandler(ctx *gin.Context) {
	var login Login
	var user model.Users
	if err := ctx.MustBindWith(&login, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    statuscode.StatusInvalidRequest,
			"message": "missing username or password",
			"data":    gin.H{},
		})
		return
	}
	if user = checkLogin(login); user == (model.Users{}) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    statuscode.StatusWrongPassword,
			"message": "wrong username or password",
			"data":    gin.H{},
		})
		return
	}
	token, refreshToken, err := generateTokenPair(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    statuscode.StatusServerError,
			"message": "server error when generating tokens",
			"data":    gin.H{},
		})
		return
	}
	err = cache.SetRefreshTokenCache(refreshToken, user.Uid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    statuscode.StatusServerError,
			"message": "server error when caching refresh token",
			"data":    gin.H{},
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    statuscode.StatusSuccess,
		"message": "login success",
		"data": gin.H{
			"access_token":  token,
			"token_type":    "Bearer",
			"expires_in":    TokenDuration.Seconds(),
			"scope":         "all",
			"refresh_token": refreshToken,
		},
	})
	return
}
func checkLogin(login Login) model.Users {
	user := model.CheckAuth(login.Username, login.Password) // delete type
	return user
}
