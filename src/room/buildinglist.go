package room

import (
	"dormitory-system/src/model"
	"dormitory-system/statuscode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BuildingListHandler(ctx *gin.Context) {
	list := model.GetBuildingList()
	ctx.JSON(http.StatusOK, gin.H{
		"code": statuscode.StatusSuccess,
		"msg":  "success",
		"data": gin.H{
			"rows": list,
		},
	})
}
