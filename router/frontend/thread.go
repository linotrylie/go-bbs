package frontend

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/http/controllers"
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
