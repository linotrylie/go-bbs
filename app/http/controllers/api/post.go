package api

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/http/model/requests"
	"go-bbs/app/http/model/response"
	"go-bbs/global"
	"go.uber.org/zap"
)

type PostController struct {
}

// CommentList 获取评论列表
func (controller *PostController) CommentList(ctx *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			global.LOG.Error(err.Error(), zap.Error(err))
		}
	}()
	var request *requests.PostList
	if err = ctx.ShouldBindQuery(request); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err = request.Validate(); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

}
