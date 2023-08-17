package backend

import (
	"GoFreeBns/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func (controller *UserController) Index(ctx *gin.Context) {
	result := map[string]interface{}{}
	global.DB.Table("bbs_user").Where("id = ?", 1).Take(&result)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello bns",
		"user":    result,
	})
}
