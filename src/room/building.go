package room

import (
	"dormitory-system/src/model"
	"dormitory-system/statuscode"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func BuildingHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    statuscode.StatusInvalidRequest,
			"message": "bad request",
			"data":    gin.H{},
		})
		return
	}
	info := model.GetBuildingInfo(id)
	ctx.JSON(http.StatusOK,
		gin.H{
			"code":    statuscode.StatusSuccess,
			"message": "success",
			"data":    info,
		})
}
