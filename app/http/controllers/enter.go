package controllers

import (
	"GoFreeBns/app/http/controllers/api"
	"GoFreeBns/app/http/controllers/backend"
	"GoFreeBns/app/http/controllers/common"
	"GoFreeBns/app/http/controllers/frontend"
)

type AllRouterGroup struct {
	ApiGroup      api.ApiGroup
	BackendGroup  backend.BackendGroup
	CommonGroup   common.CommonGroup
	FrontendGroup frontend.FrontendGroup
}

var AllRouterGroupApp = new(AllRouterGroup)
