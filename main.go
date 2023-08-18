package main

import (
	"GoFreeBns/core"
	"GoFreeBns/global"
	"GoFreeBns/initialize"
	"GoFreeBns/utils"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	global.VP = core.Viper()
	global.LOG = core.Zap()
	initialize.Redis()
	initialize.OtherInit()
	initialize.InitCrontab()
	//defer cron.Stop()
	global.DB = initialize.Gorm()
	router := initialize.Routers()
	initialize.InitViews(router) //加载模板渲染库
	srv := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    120,
		WriteTimeout:   120,
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
	}
	if utils.RunModeIsRelease() {
	}
}
