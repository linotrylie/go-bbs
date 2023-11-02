package initialize

import (
	"fmt"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-bbs/app/http/middleware"
	"go-bbs/global"
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

	Router.Use(
		middleware.Cors(),
		middleware.Recovery(true),
		middleware.RateLimitMiddleware(),
		middleware.DefaultLimit(),
		middleware.PreventDuplication(),
	)

	///////////普罗米修斯添加到中间件////////////////////
	global.RegisterPrometheus(global.Prome, "go-bbs", ":8080")
	Router.Use(global.NewPrometheusHandle(global.Prome))
	///////////普罗米修斯添加到中间件////////////////////

	apiRouter := router.AllRouterGroupMain.ApiRouterGroup
	backendRouter := router.AllRouterGroupMain.BackendRouterGroup
	commonRouter := router.AllRouterGroupMain.CommonRouterGroup

	Router.Static("/storage/uploads/file/", "./storage/uploads/file/") //静态文件目录
	//Router.LoadHTMLGlob("views/*")
	Router.GET("/", func(context *gin.Context) {})
	Router.GET("/metrics", global.PromHandler(promhttp.Handler()))
	Router.GET("/favicon.ico", func(context *gin.Context) {})
	//公共路由组件 不需要鉴权
	PublicGroup := Router.Group(global.CONFIG.System.RouterPrefix)
	{
		// 健康监测
		PublicGroup.GET("/ding", func(c *gin.Context) {
			c.JSON(http.StatusOK, "dong")
		})
	}
	{
		commonRouter.InitCaptchaRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}

	PrivateGroup := Router.Group("/")
	Router.HTMLRender = ginview.Default()
	PrivateGroup.Use(middleware.JWT(), middleware.RequestLogger())
	{
		apiRouter.InitAuthRouter(PrivateGroup)
		apiRouter.InitUserRouter(PrivateGroup, PublicGroup)      //前端用户
		apiRouter.InitForumRouter(PrivateGroup, PublicGroup)     //版块
		apiRouter.InitThreadRouter(PrivateGroup, PublicGroup)    //帖子
		apiRouter.InitPostRouter(PrivateGroup, PublicGroup)      //帖子评论，回复帖子，发表帖子等等
		apiRouter.InitKaDaoDataRouter(PrivateGroup, PublicGroup) //卡刀数据
		commonRouter.InitUploadRouter(PrivateGroup)
		mw := ginview.NewMiddleware(goview.Config{
			Root:      "resource/views/backend",
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
		backendWithoutAuth := PublicGroup.Group("backend", mw)
		backendRouter.InitUserRouter(backend)
		backendRouter.InitTestRouter(backend, backendWithoutAuth)
	}
	Session(PublicGroup)
	Session(PrivateGroup)
	return Router
}
