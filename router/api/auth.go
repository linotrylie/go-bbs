package api

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/http/controllers"
)

type AuthRouter struct {
}

func (ar *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {
	Auth := Router.Group("auth")
	authContr := controllers.AllRouterGroupApp.ApiGroup.AuthController
	{
		Auth.GET("index", authContr.Index)
	}
}
