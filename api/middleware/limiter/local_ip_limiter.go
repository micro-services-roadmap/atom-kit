package limiter

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/micro-services-roadmap/atom-kit/kg"
)

type LocalLimit struct {
	LimitBase
	localCache sync.Map
}

func (l *LocalLimit) CheckOrMark(key string, expire int, limit int) error {
	val, ok := l.localCache.LoadOrStore(key, 1)
	if !ok {
		go func() {
			time.Sleep(time.Duration(expire) * time.Second)
			l.localCache.Delete(key)
		}()
	}

	count := val.(int)
	if count > limit {
		return ErrLimited
	}

	// 次数加一
	l.localCache.Store(key, count+1)
	return nil
}

func LocalLimiter() gin.HandlerFunc {

	limiter := LocalLimit{
		localCache: sync.Map{},
	}
	limiter.Expire = kg.C.System.LimitTimeIP
	limiter.Limit = kg.C.System.LimitCountIP

	return limiter.Process(limiter.CheckOrMark, limiter.Expire, limiter.Limit)
}
