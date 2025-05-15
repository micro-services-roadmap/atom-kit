package limiter

import (
	"errors"
	"github.com/kongmsr/oneid-core/modelo"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro-services-roadmap/atom-kit/kg"
)

const (
	disableLimiter = iota
	localLimiter
	redisLimiter
)

var ErrLimited = errors.New("rate limit exceeded")

type CheckFn = func(key string, expire int, limit int) (err error)

type LimitInterface interface {
	// GenerationKey 根据业务生成key 下面CheckOrMark查询生成
	GenerationKey(c *gin.Context) string

	// CheckOrMark 检查函数,用户可修改具体逻辑,更加灵活
	CheckOrMark(key string, expire int, limit int) error
}

type LimitBase struct {
	LimitInterface
	// Expire key 过期时间
	Expire int
	// Limit 周期时间
	Limit int
}

func (l *LimitBase) GenerationKey(c *gin.Context) string {
	return "wordpress-plus_Limit" + c.ClientIP()
}

func (l *LimitBase) Process(check CheckFn, expire int, limit int) gin.HandlerFunc {

	return func(c *gin.Context) {
		if err := check(l.GenerationKey(c), expire, limit); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": modelo.ERROR, "msg": err.Error()})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}

func Limiter() gin.HandlerFunc {
	cacheType := kg.C.System.LimitType

	if disableLimiter == cacheType {
		return func(c *gin.Context) {
			c.Next()
		}
	}

	if localLimiter == cacheType {
		return LocalLimiterV2()
	}

	if redisLimiter == cacheType {
		return RedisLimiter()
	}

	panic("非法限流(缓存)类型")
}
