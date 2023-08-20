package service

import (
	"github.com/duke-git/lancet/v2/compare"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go-bbs/app/http/model/requests"
	"go-bbs/global"
	"go-bbs/utils"
	"time"
)

type CaptchaService struct {
	EmailCaptchaVerify requests.EmailCaptchaVerify
	CaptchaVerify      requests.CaptchaVerify
}

var store = base64Captcha.DefaultMemStore

func (serv *CaptchaService) VerifyCaptcha(c *gin.Context) bool {
	key := c.ClientIP()
	openCaptcha := global.CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}
	var oc = openCaptcha == 0 || openCaptcha < utils.InterfaceToInt(v)
	if !oc {
		_ = global.BlackCache.Increment(key, 1)
		return false
	}
	//判断系统是开启了哪一项的验证
	switch global.CONFIG.Captcha.IsEmailOrPic {
	case 0:
		err := verifyEmailCaptcha(serv.EmailCaptchaVerify)
		if !err {
			return false
		}
		err = verifyPicCaptcha(serv.CaptchaVerify)
		if !err {
			return false
		}
		return true
	case 1:
		return verifyPicCaptcha(serv.CaptchaVerify)
	case 2:
		return verifyEmailCaptcha(serv.EmailCaptchaVerify)
	default:
		return false
	}
}

func verifyEmailCaptcha(EmailCaptchaVerify requests.EmailCaptchaVerify) bool {
	codeInter, errBool := global.BlackCache.Get(EmailCaptchaVerify.Email)
	if !errBool {
		return false
	}
	codeStr := convertor.ToString(codeInter)
	if compare.EqualValue(EmailCaptchaVerify.Value, codeStr) {
		global.BlackCache.Delete(EmailCaptchaVerify.Email)
		return true
	}
	return false
}
func verifyPicCaptcha(CaptchaVerify requests.CaptchaVerify) bool {
	if store.Verify(CaptchaVerify.Key, CaptchaVerify.Value, true) {
		return true
	}
	return false
}
