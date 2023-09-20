package api

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/http/controllers"
)

type PostRouter struct {
}

func (tr *PostRouter) InitPostRouter(privateRouter *gin.RouterGroup, publicRouter *gin.RouterGroup) {
	postWithoutAuth := publicRouter.Group("post")
	postWithAuth := privateRouter.Group("post")
	postContr := controllers.AllRouterGroupApp.ApiGroup.PostController
	{
		postWithoutAuth.GET("comment-list", postContr.CommentList)
		postWithoutAuth.POST("create", postContr.Create)
	}
	{
		postWithAuth.POST("detail", postContr.CommentList)
	}
}
