package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"go-bbs/global"
	"net/http"
	"time"
)

var limiter = ratelimit.NewBucketWithQuantum(time.Second, 100, 100)

func RateLimitMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		if limiter.TakeAvailable(1) == 0 {
			global.LOG.Error(fmt.Sprintf("available token :%d", limiter.Available()))
			context.AbortWithStatusJSON(http.StatusTooManyRequests, "Too Many Request")
		} else {
			context.Writer.Header().Set("X-RateLimit-Remaining", fmt.Sprintf("%d", limiter.Available()))
			context.Writer.Header().Set("X-RateLimit-Limit", fmt.Sprintf("%d", limiter.Capacity()))
			context.Next()
		}
	}
}
