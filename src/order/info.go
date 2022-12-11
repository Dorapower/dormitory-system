package order

import (
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func InfoHandler(ctx *gin.Context) {
	//get order info
	orderId, err := strconv.Atoi(ctx.Query("order_id"))
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    1,
			"message": "bad request",
			"data":    gin.H{},
		})
		return
	}
	info := model.GetOrderInfo(orderId)
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": info,
	})
}
