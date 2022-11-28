package room

import (
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
)

func BuildingListHandler(ctx *gin.Context) {
	list := model.GetBuildingList()
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"rows": list,
		},
	})
}
