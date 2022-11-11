package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RefreshHandler(ctx *gin.Context) {
	// TODO: refresh token
	ctx.JSON(http.StatusOK, gin.H{})
}
