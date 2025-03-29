package limiter

import (
	"context"
	"github.com/kongmsr/oneid-core/modelo"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Allower rate.NewLimiter(rate.Every(time.Minute), 1)
type Allower interface {
	Allow() bool
}

// NewErrorLimiter 达到限流阈值后直接报错
func NewErrorLimiter(limit Allower) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limit.Allow() {
			c.JSON(http.StatusOK, gin.H{"code": modelo.ERROR, "msg": ErrLimited.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}

type Waiter interface {
	Wait(ctx context.Context) error
}

// NewDelayLimiter 达到限流阈值后持有相关请求并等到执行
func NewDelayLimiter(limit Waiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := limit.Wait(c); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": modelo.ERROR, "msg": err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}
