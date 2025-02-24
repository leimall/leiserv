package global

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"

	"leiserv/utils/timer"

	"github.com/songzhibin97/gkit/cache/local_cache"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"leiserv/config"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	MALL_DB                  *gorm.DB
	MALL_DBList              map[string]*gorm.DB
	MALL_REDIS               redis.UniversalClient
	MALL_MONGO               *qmgo.QmgoClient
	MALL_CONFIG              config.Server
	MALL_VP                  *viper.Viper
	MALL_LOG                 *zap.Logger
	MALL_Timer               timer.Timer = timer.NewTimerTask()
	MALL_Concurrency_Control             = &singleflight.Group{}
	MALL_ROUTERS             gin.RoutesInfo
	BlackCache               local_cache.Cache
	lock                     sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return MALL_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := MALL_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
