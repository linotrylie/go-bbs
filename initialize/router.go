package initialize

import (
	"GoFreeBns/app/http/middleware"
	"GoFreeBns/router"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	InstallPlugin(Router)

	Router.Use(gin.Logger(), middleware.Cors(), middleware.Recovery())

	apiRouter := router.AllRouterGroupMain.ApiRouterGroup
	backendRouter := router.AllRouterGroupMain.BackendRouterGroup
	commonRouter := router.AllRouterGroupMain.CommonRouterGroup
	frontendRouter := router.AllRouterGroupMain.FrontendRouterGroup

	Router.GET("/", func(context *gin.Context) {
	})
	Router.GET("/favicon.ico", func(context *gin.Context) {
	})
	//公共路由组件 不需要鉴权
	PublicGroup := Router.Group("/")
	{
		// 健康监测
		PublicGroup.GET("/dong", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		commonRouter.InitCommonRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		frontendRouter.InitThreadRouter(PublicGroup)
		frontendRouter.InitHomeRouter(PublicGroup)
	}

	PrivateGroup := Router.Group("/s")
	{
		api := PrivateGroup.Group("api")
		apiRouter.InitAuthRouter(api)

		mw := ginview.NewMiddleware(goview.Config{
			Root:      "views/backend",
			Extension: ".html",
			Master:    "layouts/master",
			Partials:  []string{},
			Funcs: template.FuncMap{
				"copy": func() string {
					return time.Now().Format("2006")
				},
			},
			DisableCache: true,
		})
		backend := PrivateGroup.Group("backend", mw)
		backendRouter.InitUserRouter(backend)
	}

	return Router
}
