package api

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"go-bbs/app/exceptions"
	"go-bbs/app/http/model/response"
	"go-bbs/global"
	"go.uber.org/zap"
	"math"
	"strconv"
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
	forumId, _ := strconv.Atoi(ctx.DefaultQuery("fid", "1"))
	err = validation.Validate(forumId,
		validation.Min(1).Error(exceptions.ParamInvalid.Error()),
		validation.Required.Error(exceptions.ParamInvalid.Error()),
		validation.Max(math.MaxInt).Error(exceptions.ParamInvalid.Error()),
	)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	threadId, _ := strconv.Atoi(ctx.DefaultQuery("tid", "1"))
	err = validation.Validate(threadId,
		validation.Required.Error(exceptions.ParamInvalid.Error()),
		validation.Min(1).Error(exceptions.ParamInvalid.Error()),
		validation.Max(math.MaxInt).Error(exceptions.ParamInvalid.Error()),
	)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	threadVo, postVo, CommentList, err := threadService.Detail(forumId, threadId)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	result := map[string]interface{}{
		"thread":  threadVo,
		"post":    postVo,
		"comment": CommentList,
	}
	response.OkWithData(result, ctx)
	return
}

func (controller *ThreadController) name() {

}
