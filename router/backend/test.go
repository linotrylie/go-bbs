package backend

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/http/controllers"
)

type TestRouter struct {
}

func (ur *TestRouter) InitTestRouter(Router *gin.RouterGroup, publicRouter *gin.RouterGroup) {
	_ = Router.Group("test")
	TestRouterWithoutAuth := publicRouter.Group("test")
	testContro := controllers.AllRouterGroupApp.BackendGroup.TestController
	{
		TestRouterWithoutAuth.GET("/editor", testContro.Editor)
	}
}
