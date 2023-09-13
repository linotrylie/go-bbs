package common

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/exceptions"
	"go-bbs/app/http/model"
	"go-bbs/app/http/model/requests"
	"go-bbs/app/http/model/response"
	"go-bbs/app/transform"
	"go-bbs/global"
	"go.uber.org/zap"
)

type UserController struct {
}

func (controller *UserController) name(ctx *gin.Context) {

}

func (controller *UserController) Login(ctx *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			global.LOG.Error(err.Error(), zap.Error(err))
		}
	}()
	var captcha = &requests.CaptchaVerify{}
	var userLogin = &requests.UserLogin{}
	err = ctx.ShouldBind(userLogin)
	if err != nil {
		response.FailWithMessage(exceptions.ParamInvalid.Error(), ctx)
		return
	}
	err = userLogin.Validate()
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	captcha = &userLogin.CaptchaVerify
	err = captcha.Validate()
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	//验证验证码
	/*ok := captchaService.VerifyCaptcha(ctx, captcha, &requests.EmailCaptchaVerify{})
	if !ok {
		response.FailWithMessage(exceptions.FailedVerify.Error(), ctx)
		return
	}*/
	var user = &model.User{Username: userLogin.Username, Password: userLogin.Password}
	userReturn, token, err := userService.Login(user, ctx)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithData(gin.H{
		"token": token,
		"user":  transform.TransformUser(userReturn),
	}, ctx)
	return
}
