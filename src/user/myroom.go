package user

import (
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MyRoomHandler(ctx *gin.Context) {
	uid := ctx.Keys["uid"].(int)
	room, roommate := model.GetMyRoomByUid(uid)
	ctx.JSON(http.StatusOK,
		gin.H{
			"code": 200,
			"msg":  "success",
			"data": gin.H{
				"roomName": room,
				"member": gin.H{
					"rows": roommate,
				},
			},
		})
}
