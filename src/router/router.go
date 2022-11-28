package router

import (
	"dormitory-system/src/addData"
	"dormitory-system/src/auth"
	"dormitory-system/src/middleware"
	"dormitory-system/src/room"
	"dormitory-system/src/user"
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

	protected.GET("/user/myinfo", user.MyInfoHandler)
	protected.GET("/user/myroom", user.MyRoomHandler)
	protected.POST("/user/passwd", user.PasswdHandler)

	protected.GET("/room/buildinglist", room.BuildingListHandler)
	protected.GET("/room/building", room.BuildingHandler)
	protected.GET("/room/room/:id", room.RoomHandler)
	protected.GET("/room/empty", room.EmptyHandler)

	return router
}
