package main

import (
	"dormitory-system/auth"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
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

	api := router.Group("/api")
	api.POST("/auth/login", authMiddleware.LoginHandler)
	api.GET("/auth/refresh_token", authMiddleware.RefreshHandler)

	err = router.Run(":8090")
	if err != nil {
		log.Fatal("Server Error: " + err.Error())
	}
}
