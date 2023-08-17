package middleware

import (
	"github.com/gin-gonic/gin"
)

// Recovery 错误恢复
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Abort()
			}
		}()
		c.Next()
	}
}
