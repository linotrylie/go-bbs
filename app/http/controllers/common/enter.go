package common

import (
	"go-bbs/app/http/controllers/api"
	"go-bbs/app/service"
	emailServ "go-bbs/plugin/email/service"
)

type CommonGroup struct {
	CaptchaController
	api.UserController
}

var emailService = emailServ.ServiceGroupApp.EmailService
var userService = service.ServiceGroupApp.UserService
var captchaService = service.ServiceGroupApp.CaptchaService
