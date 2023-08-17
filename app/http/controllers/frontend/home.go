package frontend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeController struct {
}

func (h *HomeController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H{
		"title": "free-bns",
		"add": func(a int, b int) int {
			return a + b
		},
	})
}
