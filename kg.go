package kit

import (
	"github.com/micro-services-roadmap/atom-kit/gormx"
	"github.com/micro-services-roadmap/atom-kit/viperx"
	"github.com/micro-services-roadmap/atom-kit/zapx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitV0(conf any, path ...string) {
	viperx.InitViperWithConf(conf, path...)
	zapx.InitZap()
	gormx.InitDB()
}

func Init(path ...string) (v *viper.Viper, l *zap.Logger, d *gorm.DB) {
	v = viperx.InitViper(path...)
	l = zapx.InitZap()
	d = gormx.InitDB()

	return
}

func InitWithConf(conf any, path ...string) (v *viper.Viper, l *zap.Logger, d *gorm.DB) {
	v = viperx.InitViperWithConf(conf, path...)
	l = zapx.InitZap()
	d = gormx.InitDB()

	return
}
