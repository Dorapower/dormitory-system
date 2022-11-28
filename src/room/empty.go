package room

import (
	"dormitory-system/src/model"
	"github.com/gin-gonic/gin"
)

func EmptyHandler(ctx *gin.Context) {
	//get empty bed info by gender
	uid := ctx.Keys["uid"].(int)
	gender := model.GetUserByUid(uid).Gender
	list := model.GetEmptyBeds(gender)
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"rows": list,
		},
	})
}
