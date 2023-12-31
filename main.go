package main

import (
	"context"
	"fmt"
	"go-bbs/core"
	"go-bbs/global"
	"go-bbs/initialize"
	"go-bbs/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	global.VP = core.Viper()
	global.Prome = global.NewPrometheus()
	global.LOG = core.Zap()
	global.DB = initialize.Gorm()
	if global.DB == nil {
		global.LOG.Fatal("DB dont work!")
		return
	}
	initialize.Redis()
	initialize.OtherInit()
	initialize.InitCrontab()
	router := initialize.Routers()
	initialize.InitViews(router) //加载模板渲染库
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", global.CONFIG.System.Addr),
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			msg := fmt.Sprintf("ListenAndServe err: %v", err)
			global.LOG.Fatal(msg)
		}
	}()
	// 优雅重启
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	if err := srv.Shutdown(ctx); err != nil {
		global.LOG.Fatal("server shutdown!")
	}
	if utils.RunModeIsRelease() {
	}
}
