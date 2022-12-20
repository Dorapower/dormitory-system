package order

import (
	"dormitory-system/src/model"
	"dormitory-system/statuscode"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CreateRequest struct {
	GroupId    *int `json:"group_id" binding:"required"`
	BuildingId int  `json:"building_id" binding:"required"`
}

func CreateHandler(ctx *gin.Context) {
	//create order
	var request CreateRequest
	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    statuscode.StatusInvalidRequest,
			"message": "bad request",
			"data":    gin.H{},
		})
		return
	}
	uid := ctx.Keys["uid"].(int)
	/*
		msg, err := json.Marshal(request)
		if err != nil {
			ctx.JSON(500, gin.H{
				"code":    status.StatusServerError,
				"message": "internal server error",
				"data":    gin.H{},
			})
			return
		}
		rabbitmq.PublishOrderMessage(msg)
	*/
	orderId := model.CreateOrder(uid, *request.GroupId, request.BuildingId, int(time.Now().Unix()))
	if orderId == -1 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    statuscode.StatusCreateOrderFailed,
			"message": "create order failed",
			"data":    gin.H{},
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    statuscode.StatusSuccess,
		"message": "create order success",
		"data": gin.H{
			"order_id": orderId,
		},
	})
}
