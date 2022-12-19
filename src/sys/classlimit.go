package sys

import (
	"dormitory-system/src/model"
	"dormitory-system/statuscode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ClassLimitHandler(ctx *gin.Context) {
	classLimit := model.GetSystemConfigByKey("class_limit").KeyValue
	ctx.JSON(http.StatusOK, gin.H{
		"code": statuscode.StatusSuccess,
		"msg":  "success",
		"data": gin.H{
			"class_limit": classLimit,
		},
	})
}
