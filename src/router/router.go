package router

import (
	"dormitory-system/src/addData"
	"dormitory-system/src/auth"
	"dormitory-system/src/middleware"
	"dormitory-system/src/order"
	"dormitory-system/src/room"
	"dormitory-system/src/sys"
	"dormitory-system/src/team"
	"dormitory-system/src/user"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	api := router.Group("/")
	api.POST("/auth/login", auth.LoginHandler)
	api.POST("/auth/refresh_token", auth.RefreshHandler)

	add := router.Group("/add")
	add.POST("/auth", addData.AddAuth)
	add.POST("/user", addData.AddUser)

	protected := router.Group("/")
	protected.Use(middleware.JwtAuth())
	protected.POST("/auth/logout", auth.LogoutHandler)

	protected.GET("/user/myinfo", user.MyInfoHandler)
	protected.GET("/user/myroom", user.MyRoomHandler)
	protected.POST("/user/passwd", user.PasswdHandler)

	protected.GET("/room/buildinglist", room.BuildingListHandler)
	protected.GET("/room/building", room.BuildingHandler)
	protected.GET("/room/room", room.RoomHandler)
	protected.GET("/room/empty", room.EmptyHandler)

	protected.POST("/team/create", team.CreateTeamHandler)
	protected.POST("/team/del", team.DeleteTeamHandler)
	protected.POST("/team/join", team.JoinTeamHandler)
	protected.POST("/team/quit", team.LeaveTeamHandler)
	protected.GET("/team/my", team.MyTeamHandler)
	protected.POST("/team/transfer", team.TransferTeamHandler)

	protected.POST("/order/create", order.CreateHandler)
	protected.GET("/order/list", order.ListHandler)
	protected.GET("/order/info", order.InfoHandler)

	protected.GET("/sys/opentime", sys.OpentimeHandler)
	protected.GET("/sys/groupnum", sys.GroupNumHandler)
	protected.GET("/sys/classlimit", sys.ClassLimitHandler)

	return router
}
