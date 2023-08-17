package api

import (
	"GoFreeBns/app/http/controllers"
	"github.com/gin-gonic/gin"
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
