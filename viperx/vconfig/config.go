package vconfig

import (
	"github.com/kongmsr/oneid-core/modelo"
	"github.com/micro-services-roadmap/kit-common/api/aconfig"
	"github.com/micro-services-roadmap/kit-common/gormx/gconfig"
	"github.com/micro-services-roadmap/kit-common/redis/rconfig"
	"github.com/micro-services-roadmap/kit-common/zapx/config"
)

type Server struct {
	System System     `mapstructure:"system" json:"system" yaml:"system"`
	Email  Email      `mapstructure:"email" json:"email" yaml:"email"`
	Zap    config.Zap `mapstructure:"zap" json:"zap" yaml:"zap"`

	gconfig.DbServer `yaml:",inline" mapstructure:",squash"`
	Redis            rconfig.Redis `mapstructure:"redis" json:"redis" yaml:"redis"`

	// 跨域配置
	Cors      aconfig.CORS     `mapstructure:"cors" json:"cors" yaml:"cors"`
	JWT       aconfig.JWT      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Captcha   aconfig.Captcha  `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	OneidConf modelo.OneidConf `mapstructure:"oneid-token" json:"oneid-token" yaml:"oneid-token"`
}

type System struct {
	gconfig.DbSys `yaml:",inline" mapstructure:",squash"`

	Env          string `mapstructure:"env" json:"env" yaml:"env"`                      // 环境值
	JasyptPwd    string `mapstructure:"jasypt-pwd" json:"jasypt-pwd" yaml:"jasypt-pwd"` // 配置加解密值
	Addr         int    `mapstructure:"addr" json:"addr" yaml:"addr"`                   // 端口值
	UseRedis     bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`    // 使用redis
	LimitType    int    `mapstructure:"limit-type" json:"limit-type" yaml:"limit-type"` // 限流类型:0(本地)|1(本地)|2(redis)
	LimitCountIP int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP  int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
	OssType      string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"` // Oss类型
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	IPQueryKey   string `mapstructure:"ip-query-key" json:"ip-query-key" yaml:"ip-query-key"`
	CurrencyKey  string `mapstructure:"currency-key" json:"currency-key" yaml:"currency-key"`
}
