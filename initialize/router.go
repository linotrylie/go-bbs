package initialize

import (
	"fmt"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-bbs/app/http/middleware"
	"go-bbs/router"
	"html/template"
	"net/http"
	"time"
)

func Routers() *gin.Engine {
	Router := gin.New()
	InstallPlugin(Router)
	Router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.DateTime),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	Router.Use(middleware.Cors(), middleware.Recovery())

	///////////普罗米修斯添加到中间件////////////////////
	var p = &Prometheus{}
	registerPrometheus(p, "go-bbs", ":8080")
	Router.Use(newPrometheusHandle(p))
	///////////普罗米修斯添加到中间件////////////////////

	apiRouter := router.AllRouterGroupMain.ApiRouterGroup
	backendRouter := router.AllRouterGroupMain.BackendRouterGroup
	commonRouter := router.AllRouterGroupMain.CommonRouterGroup
	frontendRouter := router.AllRouterGroupMain.FrontendRouterGroup

	Router.GET("/", func(context *gin.Context) {
	})

	Router.GET("/metrics", PromHandler(promhttp.Handler()))

	Router.GET("/favicon.ico", func(context *gin.Context) {
	})
	//公共路由组件 不需要鉴权
	PublicGroup := Router.Group("/")
	{
		// 健康监测
		PublicGroup.GET("/ding", func(c *gin.Context) {
			c.JSON(http.StatusOK, "dong")
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
