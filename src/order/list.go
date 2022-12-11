package order

import (
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
)

func ListHandler(ctx *gin.Context) {
	uid := ctx.Keys["uid"].(int)
	orderList := model.GetOrderList(uid)
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"rows": orderList,
		},
	})
}
