package user

import (
	"dormitory-system/src/model"
	"dormitory-system/statuscode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MyRoomHandler(ctx *gin.Context) {
	uid := ctx.Keys["uid"].(int)
	roomId, roomName, roommate := model.GetMyRoomByUid(uid)
	if roomId == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    statuscode.StatusNoRoom,
			"message": "you don't have a room yet",
			"data":    gin.H{},
		})
	} else {
		ctx.JSON(http.StatusOK,
			gin.H{
				"code":    statuscode.StatusSuccess,
				"message": "success",
				"data": gin.H{
					"roomName": roomName,
					"roomId":   roomId,
					"member": gin.H{
						"rows": roommate,
					},
				},
			})
	}
}
