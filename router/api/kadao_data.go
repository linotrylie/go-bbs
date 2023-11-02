package api

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/http/controllers"
)

type KaDaoDataRouter struct {
}

func (tr *KaDaoDataRouter) InitKaDaoDataRouter(privateRouter *gin.RouterGroup, publicRouter *gin.RouterGroup) {
	//KaDaoDataWithoutAuth := publicRouter.Group("kadao_data")
	KaDaoDataWithAuth := privateRouter.Group("kadao-data")
	kaDaoContr := controllers.AllRouterGroupApp.ApiGroup.KaDaoDataController
	{
		//KaDaoDataWithoutAuth.POST("login", kaDaoContr.Login)
		//KaDaoDataWithoutAuth.POST("kadao-user-login", kaDaoContr)
		//KaDaoDataWithoutAuth.POST("register", kaDaoContr.Register)
	}
	{
		//KaDaoDataWithAuth.POST("detail/:uid", kaDaoContr.Detail)
		KaDaoDataWithAuth.POST("get-my-kadao-data", kaDaoContr.GetMyKaDaoData)
	}
}
