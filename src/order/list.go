package order

import (
	"dormitory-system/src/model"
	"dormitory-system/statuscode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListHandler(ctx *gin.Context) {
	uid := ctx.Keys["uid"].(int)
	orderList := model.GetOrderList(uid)
	ctx.JSON(http.StatusOK, gin.H{
		"code": statuscode.StatusSuccess,
		"msg":  "success",
		"data": gin.H{
			"rows": orderList,
		},
	})
}
