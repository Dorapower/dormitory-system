package router

import (
	"github.com/gin-gonic/gin"
	"log"
)

func InitRouter() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	api := router.Group("/api")
	api.POST("/auth/login", authMiddleware.LoginHandler)
	api.GET("/auth/refresh_token", authMiddleware.RefreshHandler)

	err := router.Run(":8090")
	if err != nil {
		log.Fatal("Server Error: " + err.Error())
	}
}
