package service

type ServiceGroup struct {
	CaptchaService CaptchaService
	ThreadService  ThreadService
	JwtService     JwtService
	UserService    UserService
	GroupService   GroupService
	UploadService  UploadService
	ForumService   ForumService
}

var ServiceGroupApp = new(ServiceGroup)
