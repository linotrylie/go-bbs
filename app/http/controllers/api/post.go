package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-bbs/app/http/model/requests"
	"go-bbs/app/http/model/response"
	"go-bbs/global"
	"go.uber.org/zap"
)

type PostController struct {
}

// CommentList 获取评论列表 生成注释
func (controller *PostController) CommentList(ctx *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			global.LOG.Error(err.Error(), zap.Error(err))
		}
	}()
	var request requests.PostList
	if err = ctx.ShouldBindQuery(&request); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err = request.Validate(); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	postVoList, totalPage, err := postService.CommentList(request.Tid, request.Page, request.PageSize, request.Order, request.Sort)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithData(response.PageResult{
		Page:     request.Page,
		PageSize: request.PageSize,
		Total:    totalPage,
		List:     postVoList,
	}, ctx)
	return
}

func (controller *PostController) Create(ctx *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			global.LOG.Error(err.Error(), zap.Error(err))
		}
	}()
	content := ctx.PostForm("content")
	fmt.Println(content)
}
