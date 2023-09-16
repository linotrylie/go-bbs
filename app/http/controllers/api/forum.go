package api

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/http/model/response"
	"go-bbs/global"
	"go.uber.org/zap"
	"strconv"
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
	forumId, _ := strconv.Atoi(ctx.Query("forum"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	pageSize, _ := strconv.Atoi(ctx.Query("page_size"))
	list, err := forumService.ThreadList(forumId, page, pageSize)
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
