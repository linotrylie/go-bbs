package api

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/http/model/requests"
	"go-bbs/app/http/model/response"
	"go-bbs/global"
	"go.uber.org/zap"
)

type ForumController struct {
}

func (controller *ForumController) ThreadList(ctx *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			global.LOG.Error(err.Error(), zap.Error(err))
		}
	}()
	var threadList requests.ThreadList
	if err = ctx.ShouldBindQuery(&threadList); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err = threadList.Validate(); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	list, err := forumService.ThreadList(threadList.Fid, threadList.Page, threadList.PageSize, threadList.Order, threadList.Sort)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithData(list, ctx)
	return
}
func (controller *ForumController) List(ctx *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			global.LOG.Error(err.Error(), zap.Error(err))
		}
	}()
	list, err := forumService.List()
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithData(list, ctx)
	return
}
func (controller *ForumController) Detail(ctx *gin.Context) {
	ctx.JSON(0, gin.H{
		"msg": "这里是测试",
	})
	return
}
