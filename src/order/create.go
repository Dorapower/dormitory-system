package order

import (
	"dormitory-system/src/model"
	"dormitory-system/src/rabbitmq"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

type CreateRequest struct {
	GroupId    *int `json:"group_id" binding:"required"`
	BuildingId int  `json:"building_id" binding:"required"`
}

func CreateHandler(ctx *gin.Context) {
	//create order
	var createRequest CreateRequest
	if err := ctx.Bind(&createRequest); err != nil {
		ctx.JSON(400, gin.H{
			"code":    1,
			"message": "bad request",
			"data":    gin.H{},
		})
		return
	}
	uid := ctx.Keys["uid"].(int)
	msg, err := json.Marshal(createRequest)
	if err != nil {
		ctx.JSON(500, gin.H{
			"code":    2,
			"message": "server error",
			"data":    gin.H{},
		})
		return
	}
	rabbitmq.PublishOrderMessage(msg)
	orderId := model.CreateOrder(uid, *createRequest.GroupId, createRequest.BuildingId, int(time.Now().Unix()))
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "create order success",
		"data": gin.H{
			"order_id": orderId,
		},
	})
}
