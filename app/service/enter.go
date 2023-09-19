package service

type ServiceGroup struct {
	CaptchaService captchaService
	ThreadService  threadService
	JwtService     jwtService
	UserService    userService
	GroupService   groupService
	UploadService  uploadService
	ForumService   forumService
	PostService    postService
}

var ServiceGroupApp = new(ServiceGroup)
