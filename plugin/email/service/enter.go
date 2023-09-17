package service

type ServiceGroup struct {
	EmailService emailService
}

var ServiceGroupApp = new(ServiceGroup)
