package common

import (
	"go-bbs/app/service"
	emailServ "go-bbs/plugin/email/service"
)

type CommonGroup struct {
	CaptchaController
	UserController
}

var emailService = emailServ.ServiceGroupApp.EmailService
var userService = service.ServiceGroupApp.UserService
var captchaService = service.ServiceGroupApp.CaptchaService
