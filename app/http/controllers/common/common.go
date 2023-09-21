package common

import (
	"github.com/gin-gonic/gin"
	"go-bbs/global"
	"go.uber.org/zap"
)

type CommonController struct {
}

// GenerateFormToken 生成表单token 防止用户重复提交
func (controller *CommonController) GenerateFormToken(ctx *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			global.LOG.Error(err.Error(), zap.Error(err))
		}
	}()
}

func (controller *CommonController) name(ctx *gin.Context) {

}
