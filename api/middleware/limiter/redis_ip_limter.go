package limiter

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/micro-services-roadmap/atom-kit/kg"
)

type RedisLimit struct {
	LimitBase
	localCache sync.Map
}

func (l *RedisLimit) CheckOrMark(key string, expire int, limit int) error {
	if kg.REDIS == nil {
		return errors.New("redis is not init")
	}

	count, err := kg.REDIS.Exists(context.Background(), key).Result()
	if err != nil {
		return err
	}

	if count == 0 {
		pipe := kg.REDIS.TxPipeline()
		pipe.Incr(context.Background(), key)
		pipe.Expire(context.Background(), key, time.Duration(expire)*time.Second)
		_, err = pipe.Exec(context.Background())
		return err
	}

	// 次数
	if times, err := kg.REDIS.Get(context.Background(), key).Int(); err != nil {
		return err
	} else {
		if times < limit {
			return kg.REDIS.Incr(context.Background(), key).Err()
		}

		if t, err := kg.REDIS.PTTL(context.Background(), key).Result(); err != nil {
			return errors.New("请求太过频繁，请稍后再试")
		} else {
			return errors.New("请求太过频繁, 请 " + t.String() + " 秒后尝试")
		}
	}
}

func RedisLimiter() gin.HandlerFunc {

	limiter := RedisLimit{}
	limiter.Expire = kg.C.System.LimitTimeIP
	limiter.Limit = kg.C.System.LimitCountIP

	return limiter.Process(limiter.CheckOrMark, limiter.Expire, limiter.Limit)
}
