package api

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/exceptions"
	"go-bbs/app/http/model/requests"
	"go-bbs/app/http/model/response"
	"go-bbs/global"
	"go.uber.org/zap"
)

type KaDaoDataController struct {
}

func (controller *KaDaoDataController) GetMyKaDaoData(ctx *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			global.LOG.Error(err.Error(), zap.Error(err))
		}
	}()
	var getMyKaDaoDataRequest = &requests.GetKaDaoDataRequest{}
	err = ctx.ShouldBind(getMyKaDaoDataRequest)
	if err != nil {
		response.FailWithMessage(exceptions.ParamInvalid.Error(), ctx)
		return
	}
	list, totalPage, err := kaDaoDataService.GetKaDaoDataList(
		getMyKaDaoDataRequest.Username, getMyKaDaoDataRequest.Sort, getMyKaDaoDataRequest.Order,
		getMyKaDaoDataRequest.Page, getMyKaDaoDataRequest.PageSize)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	pageRes := response.PageResult{Page: getMyKaDaoDataRequest.Page, PageSize: getMyKaDaoDataRequest.PageSize,
		Total: totalPage, List: list}
	response.OkWithData(pageRes, ctx)
	return
}

func (controller *KaDaoDataController) name(ctx *gin.Context) {

}
