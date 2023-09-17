package api

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/http/controllers"
	"go-bbs/app/http/middleware"
)

type ThreadRouter struct {
}

func (tr *ForumRouter) InitThreadRouter(privateRouter *gin.RouterGroup, publicRouter *gin.RouterGroup) {
	ForumWithoutAuth := publicRouter.Group("thread").Use(middleware.AuthForum())
	Forum := publicRouter.Group("thread")
	ForumWithAuth := privateRouter.Group("thread").Use(middleware.AuthForum())
	threadContr := controllers.AllRouterGroupApp.ApiGroup.ThreadController
	{
		Forum.GET("thread-list", threadContr.Detail)
		ForumWithoutAuth.GET("detail", threadContr.Detail)
	}
	{
		ForumWithAuth.POST("create", threadContr.Detail)
	}
}
