package sys

import (
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
)

func ClassLimitHandler(ctx *gin.Context) {
	classLimit := model.GetSystemConfigByKey("class_limit").KeyValue
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"class_limit": classLimit,
		},
	})
}
