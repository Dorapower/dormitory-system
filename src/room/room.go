package room

import (
	"dormitory-system/src/model"
	"dormitory-system/statuscode"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RoomHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": statuscode.StatusInvalidRequest,
			"msg":  "bad request",
			"data": gin.H{},
		})
	}
	info := model.GetRoomInfoById(id)
	ctx.JSON(http.StatusOK, gin.H{
		"code": statuscode.StatusSuccess,
		"msg":  "success",
		"data": info,
	})
}
