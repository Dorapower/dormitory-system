package room

import (
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func RoomHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{
			"code": 400,
			"msg":  "bad request",
			"data": gin.H{},
		})
	}
	info := model.GetRoomInfoById(id)
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": info,
	})
}
