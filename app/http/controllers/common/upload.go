package common

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/http/model"
	"go-bbs/app/http/model/response"
	"go-bbs/global"
	"go.uber.org/zap"
)

type UploadController struct {
}

func (controller *UploadController) UploadFile(ctx *gin.Context) {
	var file model.Attach
	_, header, err := ctx.Request.FormFile("file")
	if err != nil {
		global.LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", ctx)
		return
	}
	file, err = uploadService.UploadFile(header) // 文件上传后拿到文件路径
	if err != nil {
		global.LOG.Error("修改数据库链接失败!", zap.Error(err))
		response.FailWithMessage("修改数据库链接失败", ctx)
		return
	}
	response.OkWithDetailed(file, "上传成功", ctx)
}
