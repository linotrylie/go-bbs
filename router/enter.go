package router

import (
	"go-bbs/router/api"
	"go-bbs/router/backend"
	"go-bbs/router/common"
	"go-bbs/router/frontend"
)

type AllRouterGroup struct {
	ApiRouterGroup      api.RouterGroup
	BackendRouterGroup  backend.RouterGroup
	FrontendRouterGroup frontend.RouterGroup
	CommonRouterGroup   common.RouterGroup
}

var AllRouterGroupMain = new(AllRouterGroup)
