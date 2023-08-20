package controllers

import (
	"go-bbs/app/http/controllers/api"
	"go-bbs/app/http/controllers/backend"
	"go-bbs/app/http/controllers/common"
	"go-bbs/app/http/controllers/frontend"
)

type AllRouterGroup struct {
	ApiGroup      api.ApiGroup
	BackendGroup  backend.BackendGroup
	CommonGroup   common.CommonGroup
	FrontendGroup frontend.FrontendGroup
}

var AllRouterGroupApp = new(AllRouterGroup)
