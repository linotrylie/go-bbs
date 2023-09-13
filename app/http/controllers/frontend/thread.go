package frontend

import (
	"github.com/gin-gonic/gin"
	"go-bbs/utils"
	"net/http"
)

type ThreadController struct {
}

func (controller *ThreadController) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello bns",
		"ip":      utils.Long2ip(3076426154),
	})
}
