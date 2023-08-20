package backend

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/http/controllers"
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
