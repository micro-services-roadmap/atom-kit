package kg

import (
	"github.com/micro-services-roadmap/atom-kit/viperx/vconfig"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	C     vconfig.Server
	V     *viper.Viper
	L     *zap.Logger
	DB    *gorm.DB
	REDIS *redis.Client
)

const (
	DbMysql = "mysql"
	DbPgsql = "pgsql"
)
