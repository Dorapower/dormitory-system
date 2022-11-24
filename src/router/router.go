package router

import (
	"dormitory-system/src/addData"
	"dormitory-system/src/auth"
	"dormitory-system/src/middleware"
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

	add := router.Group("/add")
	add.POST("/auth", addData.AddAuth)
	add.POST("/user", addData.AddUser)

	protected := router.Group("/api")
	protected.Use(middleware.JwtAuth())
	protected.GET("/auth/logout", auth.LogoutHandler)
	return router
}
