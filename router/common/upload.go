package common

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/http/controllers"
)

type UploadRouter struct {
}

func (com *UploadRouter) InitUploadRouter(Router *gin.RouterGroup) {
	upload := Router.Group("upload")
	uploadContr := controllers.AllRouterGroupApp.CommonGroup.UploadController
	{
		upload.POST("upload", uploadContr.UploadFile)
	}
}
