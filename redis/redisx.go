package gormx

import (
	kg "github.com/micro-services-roadmap/atom-kit/kg"
	"github.com/micro-services-roadmap/atom-kit/redis/initialize"
)

func InitRedis() {
	kg.REDIS = initialize.Redis()
}
