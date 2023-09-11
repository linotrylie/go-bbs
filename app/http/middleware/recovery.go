package middleware

import (
	"github.com/gin-gonic/gin"
	"go-bbs/global"
	"go.uber.org/zap"
	"net/http"
)

// Recovery 错误恢复
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.LOG.Fatal("我捕获到panic啦："+err.(string), zap.Error(nil))
				c.String(http.StatusOK, "我捕获到panic啦：%v", err.(string))
				return
			}
		}()
		c.Next()
	}
}
