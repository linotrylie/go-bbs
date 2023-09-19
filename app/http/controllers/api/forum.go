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
	var ForumThreadList requests.ForumThreadList
	if err = ctx.ShouldBindQuery(&ForumThreadList); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err = ForumThreadList.Validate(); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	forum, threadVoList, totalPage, err := forumService.ThreadList(ForumThreadList.Fid, ForumThreadList.Page,
		ForumThreadList.PageSize, ForumThreadList.Order, ForumThreadList.Sort)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	pageRes := response.PageResult{Page: ForumThreadList.Page, PageSize: ForumThreadList.PageSize,
		Total: totalPage, List: threadVoList}
	response.OkWithData(gin.H{
		"forum":       forum,
		"thread_list": pageRes,
	}, ctx)
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
