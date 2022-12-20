package room

import (
	"dormitory-system/src/model"
	"dormitory-system/statuscode"
	"github.com/gin-gonic/gin"
	"net/http"
)

func EmptyHandler(ctx *gin.Context) {
	//get empty bed info by gender
	uid := ctx.Keys["uid"].(int)
	gender := model.GetUserByUid(uid).Gender
	list := model.GetEmptyBeds(gender)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    statuscode.StatusSuccess,
		"message": "success",
		"data": gin.H{
			"row": list,
		},
	})
}
