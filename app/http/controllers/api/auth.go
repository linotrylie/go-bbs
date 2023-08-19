package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
}

func (controller *AuthController) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello bns",
	})
}

func (controller *AuthController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello bns",
	})
}
