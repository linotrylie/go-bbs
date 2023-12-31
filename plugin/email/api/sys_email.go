package api

import (
	"github.com/gin-gonic/gin"
	"go-bbs/app/http/model/response"
	"go-bbs/global"
	emailresponse "go-bbs/plugin/email/model/response"
	"go-bbs/plugin/email/service"
	"go.uber.org/zap"
)

type EmailApi struct{}

// EmailTest
// @Tags      System
// @Summary   发送测试邮件
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200  {string}  string  "{"success":true,"data":{},"msg":"发送成功"}"
// @Router    /email/emailTest [post]
func (s *EmailApi) EmailTest(c *gin.Context) {
	err := service.EmailService.EmailTest()
	if err != nil {
		global.LOG.Error("发送失败!", zap.Error(err))
		response.FailWithMessage("发送失败", c)
		return
	}
	response.OkWithMessage("发送成功", c)
}

// SendEmail
// @Tags      System
// @Summary   发送邮件
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param     data  body      email_response.Email  true  "发送邮件必须的参数"
// @Success   200   {string}  string                "{"success":true,"data":{},"msg":"发送成功"}"
// @Router    /email/sendEmail [post]
func (s *EmailApi) SendEmail(c *gin.Context) {
	var email emailresponse.Email
	err := c.ShouldBindJSON(&email)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.EmailService.SendEmail(email.To, email.Subject, email.Body)
	if err != nil {
		global.LOG.Error("发送失败!", zap.Error(err))
		response.FailWithMessage("发送失败", c)
		return
	}
	response.OkWithMessage("发送成功", c)
}
