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
	emailService   = emailServ.EmailService
	userService    = service.UserService
	captchaService = service.CaptchaService
	uploadService  = service.UploadService
)
