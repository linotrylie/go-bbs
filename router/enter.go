package router

import (
	"GoFreeBns/router/api"
	"GoFreeBns/router/backend"
	"GoFreeBns/router/common"
	"GoFreeBns/router/frontend"
)

type AllRouterGroup struct {
	ApiRouterGroup      api.RouterGroup
	BackendRouterGroup  backend.RouterGroup
	FrontendRouterGroup frontend.RouterGroup
	CommonRouterGroup   common.RouterGroup
}

var AllRouterGroupMain = new(AllRouterGroup)
