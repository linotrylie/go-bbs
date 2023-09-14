package controllers

import (
	"go-bbs/app/http/controllers/api"
	"go-bbs/app/http/controllers/backend"
	"go-bbs/app/http/controllers/common"
)

type AllRouterGroup struct {
	ApiGroup     api.ApiGroup
	BackendGroup backend.BackendGroup
	CommonGroup  common.CommonGroup
}

var AllRouterGroupApp = new(AllRouterGroup)
