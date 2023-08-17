package common

import (
	"GoFreeBns/app/exceptions"
	"GoFreeBns/app/http/model"
	"GoFreeBns/app/http/model/requests"
	"GoFreeBns/app/http/model/response"
	"GoFreeBns/global"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"time"
)

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
// var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

type CaptchaController struct{}

func (ca *CaptchaController) EmailCaptcha(c *gin.Context) {

}

func (ca *CaptchaController) CaptchaVerify(c *gin.Context) {
	var CaptchaVerify requests.CaptchaVerify
	err := c.ShouldBindJSON(&CaptchaVerify)
	key := c.ClientIP()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = CaptchaVerify.Validate()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	openCaptcha := global.CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}
	var oc = openCaptcha == 0 || openCaptcha < interfaceToInt(v)
	if !oc || store.Verify(CaptchaVerify.Key, CaptchaVerify.Value, true) {
		response.OkWithMessage("验证成功!", c)
		return
	}
	err = global.BlackCache.Increment(key, 1)
	if err != nil {
		return
	}
	response.FailWithMessage(exceptions.FailedVerify.Error(), c)
}

// Captcha
// @Tags      Common
// @Summary   生成验证码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=model.CaptchaResponse,msg=string}  "生成验证码,返回包括随机数id,base64,验证码长度,是否开启验证码"
// @Router    /common/captcha/pic-captcha [post]
func (ca *CaptchaController) Captcha(c *gin.Context) {
	// 判断验证码是否开启
	openCaptcha := global.CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	key := c.ClientIP()
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool
	if openCaptcha == 0 || openCaptcha < interfaceToInt(v) {
		oc = true
	}
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.CONFIG.Captcha.ImgHeight, global.CONFIG.Captcha.ImgWidth, global.CONFIG.Captcha.KeyLong, 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		global.LOG.Error("验证码获取失败!", zap.Error(err))
		response.FailWithMessage("验证码获取失败!", c)
		return
	}
	response.OkWithDetailed(model.CaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.CONFIG.Captcha.KeyLong,
		OpenCaptcha:   oc,
	}, "success", c)
	return
}

// 类型转换
func interfaceToInt(v interface{}) (i int) {
	switch v := v.(type) {
	case int:
		i = v
	default:
		i = 0
	}
	return
}
