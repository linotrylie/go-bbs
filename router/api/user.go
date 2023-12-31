package api

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/http/controllers"
)

type UserRouter struct {
}

func (tr *UserRouter) InitUserRouter(privateRouter *gin.RouterGroup, publicRouter *gin.RouterGroup) {
	UserWithoutAuth := publicRouter.Group("user")
	UserWithAuth := privateRouter.Group("user")
	userContr := controllers.AllRouterGroupApp.ApiGroup.UserController
	{
		UserWithoutAuth.POST("login", userContr.Login)
		UserWithoutAuth.POST("kadao-user-login", userContr.KaDaoUserLogin)
		UserWithoutAuth.POST("register", userContr.Register)
	}
	{
		UserWithAuth.POST("detail/:uid", userContr.Detail)
		UserWithAuth.GET("detail/:uid", userContr.Detail)
		UserWithAuth.POST("edit", userContr.Edit)
		UserWithAuth.POST("change_password", userContr.ChangePassword)
		UserWithAuth.POST("logout", userContr.Logout)
	}
}
