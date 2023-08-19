package service

type ServiceGroup struct {
	CaptchaService CaptchaService
}

var ServiceGroupApp = new(ServiceGroup)
