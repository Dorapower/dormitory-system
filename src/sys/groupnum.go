package sys

import (
	"dormitory-system/src/model"
	"dormitory-system/statuscode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GroupNumHandler(ctx *gin.Context) {
	groupLimit := model.GetSystemConfigByKey("group_limit").KeyValue
	groupNum := model.GetSystemConfigByKey("group_num").KeyValue
	ctx.JSON(http.StatusOK, gin.H{
		"code":    statuscode.StatusSuccess,
		"message": "success",
		"data": gin.H{
			"group_limit": groupLimit,
			"group_num":   groupNum,
		},
	})
}
