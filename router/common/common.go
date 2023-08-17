package common

import (
	"GoFreeBns/app/http/controllers"
	"github.com/gin-gonic/gin"
)

type CommonRouter struct {
}

func (com *CommonRouter) InitCommonRouter(Router *gin.RouterGroup) {
	captcha := Router.Group("captcha")
	captchaContr := controllers.AllRouterGroupApp.CommonGroup.CaptchaController
	{
		captcha.GET("pic-captcha", captchaContr.Captcha)
		captcha.POST("verify-captcha", captchaContr.CaptchaVerify)
		captcha.GET("email-captcha", captchaContr.EmailCaptcha)
	}
}
