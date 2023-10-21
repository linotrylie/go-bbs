package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-bbs/app/exceptions"
	"go-bbs/app/http/model/requests"
	"go-bbs/app/http/model/response"
	"go-bbs/app/transform"
	"go-bbs/global"
	"go-bbs/utils"
	"go.uber.org/zap"
)

type UserController struct {
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
	userReturn, jwtCustomClaims, token, err := userService.Login(*userLogin, ctx)
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
	userVo.Group = group
	response.OkWithData(gin.H{
		"token":      token,
		"expired_at": jwtCustomClaims.ExpiresAt.Unix(),
		"user":       userVo,
	}, ctx)
	return
}

// Detail 获取用户详情
func (controller *UserController) Detail(ctx *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			global.LOG.Error(err.Error(), zap.Error(err))
		}
	}()
	addr, err := utils.GetIpCity("220.197.189.152")
	fmt.Println("addr", addr)
	fmt.Println("err", err)
	var userDetail = &requests.UserDetail{}
	err = ctx.ShouldBindUri(userDetail)
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
		response.FailWithMessage(exceptions.NotFoundData.Error(), ctx)
		return
	}
	userVo := transform.TransformUser(detail)
	group, err := groupService.Detail(detail.Gid)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	userVo.Group = group
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

func (controller *UserController) ChangePassword(ctx *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			global.LOG.Error(err.Error(), zap.Error(err))
		}
	}()
	var userChangePassword = &requests.UserChangePassword{}
	err = ctx.ShouldBind(userChangePassword)
	if err != nil {
		response.FailWithMessage(exceptions.ParamInvalid.Error(), ctx)
		return
	}
	err = userChangePassword.Validate()
	if err != nil {
		response.FailWithMessage(exceptions.ParamInvalid.Error(), ctx)
		return
	}
	if userChangePassword.NewPassword != userChangePassword.NewPasswordVerify {
		err = exceptions.FailedVerify
		response.FailWithMessage(exceptions.FailedVerify.Error(), ctx)
		return
	}
	err = userService.ChangesPassword(userChangePassword)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithMessage("修改成功", ctx)
	return
}

func (controller *UserController) Edit(ctx *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			global.LOG.Error(err.Error(), zap.Error(err))
		}
	}()
	var userEdit = &requests.UserEdit{}
	err = ctx.ShouldBind(userEdit)
	if err != nil {
		response.FailWithMessage(exceptions.ParamInvalid.Error(), ctx)
		return
	}
	err = userEdit.Validate()
	if err != nil {
		response.FailWithMessage(exceptions.ParamInvalid.Error(), ctx)
		return
	}
	err = userService.Edit(userEdit)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithMessage("修改成功", ctx)
	return
}

func (controller *UserController) Register(ctx *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			global.LOG.Error(err.Error(), zap.Error(err))
		}
	}()
	var userRegister = &requests.UserRegister{}
	err = ctx.ShouldBind(userRegister)
	if err != nil {
		response.FailWithMessage(exceptions.ParamInvalid.Error(), ctx)
		return
	}
	err = userRegister.Validate()
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	userVo, token, expiredAt, err := userService.Register(userRegister, ctx)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithData(gin.H{
		"token":      token,
		"expired_at": expiredAt,
		"user":       userVo,
	}, ctx)
	return
}

func (controller *UserController) name(ctx *gin.Context) {

}
