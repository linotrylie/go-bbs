package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go-bbs/app/constants"
	"go-bbs/app/exceptions"
	"go-bbs/app/http/model"
	"go-bbs/app/respository"
	"go-bbs/app/service"
	"go-bbs/global"
	"strings"
)

var jwtServ = new(service.JwtService)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.CONFIG.System.Env == "develop" {
			global.User = &model.User{}
			global.User.Uid = 1
			err := respository.FindByLocation(global.User)
			if err != nil {
				c.JSON(419, gin.H{
					"msg": err.Error(),
				})
				c.Abort()
			}
			c.Next()
			return
		}
		token := c.Request.Header.Get(constants.Authorization)
		fmt.Println(token)
		if token == "" {
			c.JSON(419, gin.H{
				"msg": exceptions.FailedVerify.Error(),
			})
			c.Abort()
		}
		claims, err := jwtServ.ParseToken(strings.Trim(token, "")[7:])
		if err != nil {
			c.JSON(419, gin.H{
				"msg": exceptions.FailedVerify.Error(),
			})
			c.Abort()
		}
		global.User = &model.User{}
		global.User.Uid = claims.UID
		err = respository.FindByLocation(global.User)
		if err != nil {
			c.JSON(419, gin.H{
				"msg": err.Error(),
			})
			c.Abort()
		}

		_, err = global.REDIS.Get(context.Background(), global.User.Username).Result()
		if err != nil || err == redis.Nil {
			c.JSON(419, gin.H{
				"msg": err.Error(),
			})
			c.Abort()
		}

		c.Next()
	}
}