package order

import (
	"dormitory-system/src/model"
	"dormitory-system/statuscode"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func InfoHandler(ctx *gin.Context) {
	//get order info
	orderId, err := strconv.Atoi(ctx.Query("order_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    statuscode.StatusInvalidRequest,
			"message": "bad request",
			"data":    gin.H{},
		})
		return
	}
	info := model.GetOrderInfo(orderId)
	ctx.JSON(http.StatusOK, gin.H{
		"code": statuscode.StatusSuccess,
		"msg":  "success",
		"data": info,
	})
}
