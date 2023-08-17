package frontend

import (
	"GoFreeBns/app/http/controllers"
	"github.com/gin-gonic/gin"
)

type ThreadRouter struct {
}

func (tr *ThreadRouter) InitThreadRouter(Router *gin.RouterGroup) {
	Thread := Router.Group("thread")
	threadContr := controllers.AllRouterGroupApp.FrontendGroup.ThreadController
	{
		Thread.GET("index", threadContr.Index)
	}
}
