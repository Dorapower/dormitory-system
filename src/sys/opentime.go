package sys

import (
	"dormitory-system/src/model"
	"dormitory-system/statuscode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func OpentimeHandler(ctx *gin.Context) {
	startTime := model.GetSystemConfigByKey("start_time").KeyValue
	endTime := model.GetSystemConfigByKey("end_time").KeyValue
	ctx.JSON(http.StatusOK, gin.H{
		"code":    statuscode.StatusSuccess,
		"message": "success",
		"data": gin.H{
			"start_time": startTime,
			"end_time":   endTime,
		},
	})
}
