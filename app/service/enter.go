package service

type ServiceGroup struct {
	CaptchaService CaptchaService
	ThreadService  ThreadService
	JwtService     JwtService
	UserService    UserService
	GroupService   GroupService
	UploadService  UploadService
}

var ServiceGroupApp = new(ServiceGroup)
