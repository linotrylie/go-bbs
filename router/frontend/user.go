package frontend

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/http/controllers"
)

type UserRouter struct {
}

func (tr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	User := Router.Group("user").Use()
	userContr := controllers.AllRouterGroupApp.FrontendGroup.UserController
	{
		User.GET("index", userContr.Index)
	}
}
