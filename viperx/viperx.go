package viperx

import (
	kg "github.com/micro-services-roadmap/atom-kit/kg"
	"github.com/micro-services-roadmap/atom-kit/viperx/initialize"
	"github.com/spf13/viper"
)

func InitViperV0(path ...string) {
	// init viper
	kg.V = initialize.Viper(&kg.C, path...) // 初始化Viper
}

func InitViper(path ...string) *viper.Viper {
	// init viper
	kg.V = initialize.Viper(&kg.C, path...) // 初始化Viper

	return kg.V
}

func InitViperWithConf(conf any, path ...string) *viper.Viper {
	// init viper
	kg.V = initialize.Viper(&kg.C, path...) // 初始化Viper
	initialize.Viper(conf, path...)         // 初始化Viper

	return kg.V
}
