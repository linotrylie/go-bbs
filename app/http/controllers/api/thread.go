package api

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/http/model/requests"
	"go-bbs/app/http/model/response"
	"go-bbs/global"
	"go.uber.org/zap"
)

type ThreadController struct {
}

func (controller *ThreadController) Detail(ctx *gin.Context) {
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
	threadVo, postVo, err := threadService.Detail(threadList.Fid, threadList.Tid)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	result := map[string]interface{}{
		"thread": threadVo,
		"post":   postVo,
	}
	response.OkWithData(result, ctx)
	return
}

func (controller *ThreadController) name() {

}
