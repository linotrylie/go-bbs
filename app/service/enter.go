package service

type ServiceGroup struct {
	CaptchaService CaptchaService
	ThreadService  ThreadService
	JwtService     JwtService
	UserService    UserService
	GroupService   GroupService
}

var ServiceGroupApp = new(ServiceGroup)
