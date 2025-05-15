package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/micro-services-roadmap/atom-kit/kg"
	"golang.org/x/time/rate"
)

type LocalLimitV2 struct {
	LimitBase
	limiters map[string]*rate.Limiter
}

func (l *LocalLimitV2) CheckOrMark(key string, expire int, limit int) error {

	limiter, ok := l.limiters[key]
	if !ok {
		// todo: remove limiter if necessary
		limiter = rate.NewLimiter(rate.Every(time.Duration(expire)*time.Second), limit)
		l.limiters[key] = limiter
	}

	if !limiter.Allow() {
		return ErrLimited
	}

	return nil
}

func LocalLimiterV2() gin.HandlerFunc {

	limiter := LocalLimitV2{
		limiters: make(map[string]*rate.Limiter),
	}
	limiter.Expire = kg.C.System.LimitTimeIP
	limiter.Limit = kg.C.System.LimitCountIP

	return limiter.Process(limiter.CheckOrMark, limiter.Expire, limiter.Limit)
}
