package zapx

import (
	"fmt"
	"os"

	kg "github.com/micro-services-roadmap/atom-kit/kg"
	"github.com/micro-services-roadmap/atom-kit/zapx/config"
	"github.com/micro-services-roadmap/atom-kit/zapx/internal"
	"github.com/micro-services-roadmap/atom-kit/zapx/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Zap 获取 zapx.Logger
func Zap(config config.Zap) (logger *zap.Logger) {
	if ok, _ := util.PathExists(config.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", config.Director)
		_ = os.Mkdir(config.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores(config)
	logger = zap.New(zapcore.NewTee(cores...))

	if config.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	return logger
}

func NewZap() (logger *zap.Logger) {
	return Zap(config.Zap{
		Format: "json",
	})
}

func InitZap() *zap.Logger {
	kg.L = Zap(kg.C.Zap)
	zap.ReplaceGlobals(kg.L)
	return kg.L
}
