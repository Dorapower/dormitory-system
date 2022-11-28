package sys

import (
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
)

func OpentimeHandler(ctx *gin.Context) {
	startTime := model.GetSystemConfigByKey("start_time").KeyValue
	endTime := model.GetSystemConfigByKey("end_time").KeyValue
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"start_time": startTime,
			"end_time":   endTime,
		},
	})
}
