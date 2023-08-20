package initialize

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go-bbs/global"
)

func Session(Router *gin.RouterGroup) {
	global.Cookie = cookie.NewStore([]byte("freebns"))
	session := sessions.Sessions("freebns", global.Cookie)
	Router.Use(session)
	global.Session = sessions.Default(&gin.Context{})
}
