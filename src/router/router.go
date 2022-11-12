package router

import (
	"dormitory-system/src/auth"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	api := router.Group("/api")
	api.POST("/auth/login", auth.LoginHandler)
	api.POST("/auth/refresh_token", auth.RefreshHandler)

	return router
}
