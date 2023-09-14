package backend

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/http/controllers"
)

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	_ = Router.Group("user")
	_ = controllers.AllRouterGroupApp.BackendGroup.UserController
	{
	}
}
