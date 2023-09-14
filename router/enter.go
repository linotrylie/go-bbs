package router

import (
	"go-bbs/router/api"
	"go-bbs/router/backend"
	"go-bbs/router/common"
)

type AllRouterGroup struct {
	ApiRouterGroup     api.RouterGroup
	BackendRouterGroup backend.RouterGroup
	CommonRouterGroup  common.RouterGroup
}

var AllRouterGroupMain = new(AllRouterGroup)
