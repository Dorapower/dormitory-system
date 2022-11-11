package auth

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func checkLogin(_ Login) bool {
	// TODO: check login
	return true
}

func Authenticate(ctx *gin.Context) (user interface{}, err error) {
	// check login request and return user info
	var login Login
	if err = ctx.ShouldBind(&login); err != nil {
		return nil, jwt.ErrMissingLoginValues
	}
	if checkLogin(login) {
		return login.Username, nil
	} else {
		return nil, jwt.ErrFailedAuthentication
	}
}

func Authorize(data interface{}, _ *gin.Context) bool {
	// check user info and return whether the user is authorized
	if _, ok := data.(string); ok {
		return true
	} else {
		return false
	}
}
