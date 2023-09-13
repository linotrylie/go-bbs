package common

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/exceptions"
	"go-bbs/app/http/model"
	"go-bbs/app/http/model/requests"
	"go-bbs/app/http/model/response"
	"net/http"
)

type UserController struct {
}

func (controller *UserController) name(ctx *gin.Context) {

}

func (controller *UserController) Login(ctx *gin.Context) {
	var captcha = &requests.CaptchaVerify{}
	var userLogin = &requests.UserLogin{}
	err := ctx.ShouldBind(userLogin)
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
	err = userService.Login(user)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"capt": captcha,
		"user": userLogin,
	})
}
