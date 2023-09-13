package service

type ServiceGroup struct {
	CaptchaService CaptchaService
	ThreadService  ThreadService
	JwtService     JwtService
	UserService    UserService
}

var ServiceGroupApp = new(ServiceGroup)
