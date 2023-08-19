package middleware

import (
	"GoFreeBns/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Recovery 错误恢复
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.LOG.Fatal("我捕获到panic啦：" + err.(string))
				c.String(http.StatusOK, "我捕获到panic啦："+err.(string))
			}
		}()
		c.Next()
	}
}
