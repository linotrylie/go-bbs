package common

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/http/controllers"
)

type UserRouter struct {
}

func (tr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	UserWithoutAuth := Router.Group("user")
	userContr := controllers.AllRouterGroupApp.CommonGroup.UserController
	{
		UserWithoutAuth.POST("login", userContr.Login)
	}
}
