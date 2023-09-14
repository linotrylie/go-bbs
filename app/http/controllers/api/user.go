package api

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
	userReturn, jwtCustomClaims, token, err := userService.Login(user, ctx)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	userVo := transform.TransformUser(userReturn)
	group, err := groupService.Detail(userReturn.Gid)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	userVo.Group = *group
	response.OkWithData(gin.H{
		"token":      token,
		"expired_at": jwtCustomClaims.ExpiresAt.Second(),
		"user":       userVo,
	}, ctx)
	return
}

func (controller *UserController) Detail(ctx *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			global.LOG.Error(err.Error(), zap.Error(err))
		}
	}()
	var userDetail = &requests.UserDetail{}
	err = ctx.ShouldBind(userDetail)
	if err != nil {
		response.FailWithMessage(exceptions.ParamInvalid.Error(), ctx)
		return
	}
	err = userDetail.Validate()
	if err != nil {
		response.FailWithMessage(exceptions.ParamInvalid.Error(), ctx)
		return
	}
	detail, err := userService.Detail(userDetail.Uid)
	if err != nil {
		return
	}
	userVo := transform.TransformUser(detail)
	group, err := groupService.Detail(detail.Gid)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	userVo.Group = *group
	response.OkWithData(gin.H{
		"user": userVo,
	}, ctx)
	return
}

func (controller *UserController) Logout(ctx *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			global.LOG.Error(err.Error(), zap.Error(err))
		}
	}()
	userService.Logout()
	response.OkWithMessage("退出成功", ctx)
	return
}