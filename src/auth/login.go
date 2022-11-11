package auth

import (
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
)

func LoginHandler(ctx *gin.Context) {
	// TODO: login
	var login Login
	var user model.User
	if err := ctx.MustBindWith(&login, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error_code": 1,
			"message":    "missing username or password",
			"data":       gin.H{},
		})
	}
	if user = checkLogin(login); user == (model.User{}) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error_code": 2,
			"message":    "wrong username or password",
			"data":       gin.H{},
		})
	}
	token, refreshToken, err := generateTokenPair(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error_code": 3,
			"message":    "server error when generating tokens",
			"data":       gin.H{},
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"error_code": 0,
		"message":    "login success",
		"data": gin.H{
			"token":         token,
			"refresh_token": refreshToken,
			"expires_in":    300,
		},
	})
}
func checkLogin(login Login) model.User {
	// TODO: check login
	type_, _ := strconv.Atoi(login.Type)
	user := model.CheckAuth(login.Username, login.Password, type_)
	return user
}
