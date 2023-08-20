package frontend

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/http/controllers"
)

type HomeRouter struct {
}

func (ho *HomeRouter) InitHomeRouter(Router *gin.RouterGroup) {
	home := Router.Group("home")
	homeContr := controllers.AllRouterGroupApp.FrontendGroup.HomeController
	{
		home.GET("index", homeContr.Index)
	}
}
