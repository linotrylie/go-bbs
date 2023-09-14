package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/songzhibin97/gkit/cache/singleflight"
	"github.com/spf13/viper"
	"go-bbs/app/http/model"
	"go-bbs/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
	// 导入session存储引擎
	"github.com/gin-contrib/sessions/cookie"
)

var (
	VP         *viper.Viper
	DB         *gorm.DB
	REDIS      *redis.Client
	CONFIG     config.Server
	LOG        *zap.Logger
	Cookie     cookie.Store
	BlackCache local_cache.Cache
	Prome      *Prometheus
	Sf         = &singleflight.Group{}
	User       *model.User
)
