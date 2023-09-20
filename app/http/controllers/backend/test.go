package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestController struct {
}

func (controller *TestController) Editor(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "editor", gin.H{"title": "Editor"})
}

func (controller *TestController) name(ctx *gin.Context) {

}
