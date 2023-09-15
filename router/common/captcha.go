package common

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/http/controllers"
)

type CaptchaRouter struct {
}

func (com *CaptchaRouter) InitCaptchaRouter(Router *gin.RouterGroup) {
	captcha := Router.Group("captcha")
	captchaContr := controllers.AllRouterGroupApp.CommonGroup.CaptchaController
	{
		captcha.GET("pic-captcha", captchaContr.Captcha)
		captcha.POST("verify-pic-captcha", captchaContr.PicCaptchaVerify)
		captcha.GET("email-captcha", captchaContr.EmailCaptcha)
		captcha.POST("verify-email-captcha", captchaContr.EmailCaptchaVerify)
	}
}
