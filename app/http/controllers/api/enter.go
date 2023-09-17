package api

import (
	"go-bbs/app/service"
	emailServ "go-bbs/plugin/email/service"
)

type ApiGroup struct {
	AuthController
	UserController
	ForumController
	ThreadController
}

var (
	captchaService = service.CaptchaService
	emailService   = emailServ.EmailService
	userService    = service.UserService
	groupService   = service.GroupService
	forumService   = service.ForumService
	threadService  = service.ThreadService
)
