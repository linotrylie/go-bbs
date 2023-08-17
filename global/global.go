package global

import (
	"GoFreeBns/config"
	"github.com/gin-contrib/sessions"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	// 导入session存储引擎
	"github.com/gin-contrib/sessions/cookie"
)

var (
	DB      *gorm.DB
	REDIS   *redis.Client
	CONFIG  config.Server
	VP      *viper.Viper
	LOG     *zap.Logger
	Session sessions.Session
	Cookie  cookie.Store
	//GVA_Timer               timer.Timer = timer.NewTimerTask()
	//Concurrency_Control = &singleflight.Group{}
	BlackCache local_cache.Cache
)
