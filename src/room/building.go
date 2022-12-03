package room

import (
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func BuildingHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(400, gin.H{
			"code": 400,
			"msg":  "bad request",
			"data": gin.H{},
		})
		return
	}
	info := model.GetBuildingInfo(id)
	ctx.JSON(200,
		gin.H{
			"code": 200,
			"msg":  "success",
			"data": info,
		})
}
