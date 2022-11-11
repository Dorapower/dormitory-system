package middleware

import (
	"dormitory-system/src/api/auth"
	jwt "github.com/appleboy/gin-jwt/v2"
	"log"
)

func Token() {
	AuthMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "dormitory middleware",
		Key:         []byte("secret key"),
		IdentityKey: "username",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(string); ok {
				return jwt.MapClaims{
					"username": v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: auth.Authenticate,
		Authorizator:  auth.Authorize,
	})
	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}
}
