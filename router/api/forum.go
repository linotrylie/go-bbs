package api

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/http/controllers"
	"go-bbs/app/http/middleware"
)

type ForumRouter struct {
}

func (tr *ForumRouter) InitForumRouter(privateRouter *gin.RouterGroup, publicRouter *gin.RouterGroup) {
	ForumWithoutAuth := publicRouter.Group("forum").Use(middleware.AuthForum())
	Forum := publicRouter.Group("forum")
	ForumWithAuth := privateRouter.Group("forum").Use(middleware.AuthForum())
	forumContr := controllers.AllRouterGroupApp.ApiGroup.ForumController
	{
		ForumWithoutAuth.GET("thread-list", forumContr.ThreadList)
		Forum.GET("list", forumContr.List)
	}
	{
		ForumWithAuth.POST("detail", forumContr.Detail)
	}
}
