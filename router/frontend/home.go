package frontend

import (
	"GoFreeBns/app/http/controllers"
	"github.com/gin-gonic/gin"
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
