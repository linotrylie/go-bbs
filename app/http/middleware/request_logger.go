package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-bbs/app/constants"
	"go-bbs/app/exceptions"
	"go-bbs/app/http/model"
	"go-bbs/global"
	"go-bbs/utils"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

var respPool sync.Pool

func init() {
	respPool.New = func() interface{} {
		return make([]byte, 1024)
	}
}

// RequestLogger 记录请求日志
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte
		var userId int
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = io.ReadAll(c.Request.Body)
			if err != nil {
				global.LOG.Error("read body from request error:", zap.Error(err))
			} else {
				c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		} else {
			query := c.Request.URL.RawQuery
			query, _ = url.QueryUnescape(query)
			split := strings.Split(query, "&")
			m := make(map[string]string)
			for _, v := range split {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			body, _ = json.Marshal(&m)
		}
		token := c.Request.Header.Get(constants.Authorization)
		if token == "" {
			c.JSON(419, gin.H{
				"code": 7,
				"msg":  exceptions.LogBackIn.Error(),
			})
			c.Abort()
			return
		}
		claims, err := jwtServ.ParseToken(strings.Trim(token, "")[7:])
		if err != nil {
			c.JSON(419, gin.H{
				"code": 7,
				"msg":  exceptions.LogBackIn.Error(),
			})
			c.Abort()
			return
		}
		userId = claims.UID
		record := model.OperationLog{
			Ip:     utils.Ip2long(c.ClientIP()),
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
			Agent:  c.Request.UserAgent(),
			Body:   string(body),
			Uid:    userId,
		}
		// 上传文件时候 中间件日志进行裁断操作
		if strings.Contains(c.GetHeader("Content-Type"), "multipart/form-data") {
			if len(record.Body) > 1024 {
				// 截断
				newBody := respPool.Get().([]byte)
				copy(newBody, record.Body)
				record.Body = string(newBody)
				defer respPool.Put(newBody[:0])
			}
		}

		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()

		c.Next()

		latency := time.Since(now)
		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = c.Writer.Status()
		record.Latency = strconv.FormatInt(latency.Milliseconds(), 10) + "ms"
		record.Resp = writer.body.String()

		if strings.Contains(c.Writer.Header().Get("Pragma"), "public") ||
			strings.Contains(c.Writer.Header().Get("Expires"), "0") ||
			strings.Contains(c.Writer.Header().Get("Cache-Control"), "must-revalidate, post-check=0, pre-check=0") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/force-download") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/octet-stream") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/vnd.ms-excel") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/download") ||
			strings.Contains(c.Writer.Header().Get("Content-Disposition"), "attachment") ||
			strings.Contains(c.Writer.Header().Get("Content-Transfer-Encoding"), "binary") {
			if len(record.Resp) > 1024 {
				// 截断
				newBody := respPool.Get().([]byte)
				copy(newBody, record.Resp)
				record.Resp = string(newBody)
				defer respPool.Put(newBody[:0])
			}
		}
		//压入redis List
		data, _ := json.Marshal(record)
		global.REDIS.RPush(context.Background(), constants.OperationLog, string(data))
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
