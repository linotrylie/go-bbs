package frontend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ThreadController struct {
}

func (controller *ThreadController) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello bns",
	})
}
