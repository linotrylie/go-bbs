package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-bbs/utils"
	"strconv"
)

// PreventDuplication 防止重复提交
// TODO 对于提交的请求，做一个3秒的锁定处理
func PreventDuplication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		path := ctx.Request.URL.Path
		host := strconv.Itoa(int(utils.Ip2long(ctx.ClientIP())))
		fmt.Println(method, path, host)
	}
}
