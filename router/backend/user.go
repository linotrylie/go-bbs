package backend

import (
	"GoFreeBns/app/http/controllers"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	User := Router.Group("user")
	userContr := controllers.AllRouterGroupApp.BackendGroup.UserController
	{
		User.GET("index", userContr.Index)
	}
}
