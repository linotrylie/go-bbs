package common

import (
	"go-bbs/app/service"
	emailServ "go-bbs/plugin/email/service"
)

type CommonGroup struct {
	CaptchaController
	UploadController
}

var (
	emailService   = emailServ.ServiceGroupApp.EmailService
	userService    = service.ServiceGroupApp.UserService
	captchaService = service.ServiceGroupApp.CaptchaService
	uploadService  = service.ServiceGroupApp.UploadService
)
