package user

import (
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MyRoomHandler(ctx *gin.Context) {
	uid := ctx.Keys["uid"].(int)
	roomId, roomName, roommate := model.GetMyRoomByUid(uid)
	if roomId == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "you don't have a room yet",
			"data":    gin.H{},
		})
	} else {
		ctx.JSON(http.StatusOK,
			gin.H{
				"code":    200,
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
