package common

import emailServ "GoFreeBns/plugin/email/service"

type CommonGroup struct {
	CaptchaController
}

var emailService = emailServ.ServiceGroupApp.EmailService
