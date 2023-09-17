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

type ForumController struct {
}

func (controller *ForumController) ThreadList(ctx *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			global.LOG.Error(err.Error(), zap.Error(err))
		}
	}()
	forumId, _ := strconv.Atoi(ctx.DefaultQuery("fid", "0"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))
	order := ctx.DefaultQuery("order", "create_date")
	sort := ctx.DefaultQuery("sort", "desc")
	err = validation.Validate(forumId,
		validation.Min(0).Error(exceptions.ParamInvalid.Error()),
		validation.Max(math.MaxInt).Error(exceptions.ParamInvalid.Error()),
	)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err = validation.Validate(page,
		validation.Required.Error(exceptions.ParamInvalid.Error()),
		validation.Min(1).Error(exceptions.ParamInvalid.Error()),
		validation.Max(math.MaxInt).Error(exceptions.ParamInvalid.Error()),
	)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err = validation.Validate(pageSize,
		validation.Required.Error(exceptions.ParamInvalid.Error()),
		validation.Min(1).Error(exceptions.ParamInvalid.Error()),
		validation.Max(math.MaxInt).Error(exceptions.ParamInvalid.Error()),
	)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err = validation.Validate(order,
		validation.Required.Error(exceptions.ParamInvalid.Error()),
		validation.In("create_date", "last_date", "posts", "views").Error(exceptions.ParamInvalid.Error()),
	)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	err = validation.Validate(sort,
		validation.Required.Error(exceptions.ParamInvalid.Error()),
		validation.In("desc", "asc").Error(exceptions.ParamInvalid.Error()),
	)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	list, err := forumService.ThreadList(forumId, page, pageSize, order, sort)
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
