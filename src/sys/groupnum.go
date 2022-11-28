package sys

import (
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
)

func GroupNumHandler(ctx *gin.Context) {
	groupLimit := model.GetSystemConfigByKey("group_limit").KeyValue
	groupNum := model.GetSystemConfigByKey("group_num").KeyValue
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"group_limit": groupLimit,
			"group_num":   groupNum,
		},
	})
}
