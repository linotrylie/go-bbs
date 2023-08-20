package common

import emailServ "go-bbs/plugin/email/service"

type CommonGroup struct {
	CaptchaController
}

var emailService = emailServ.ServiceGroupApp.EmailService
