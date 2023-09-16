package api

import (
	"go-bbs/app/service"
	emailServ "go-bbs/plugin/email/service"
)

type ApiGroup struct {
	AuthController
	UserController
	ForumController
}

var (
	captchaService = new(service.CaptchaService)
	emailService   = new(emailServ.EmailService)
	userService    = new(service.UserService)
	groupService   = new(service.GroupService)
	forumService   = new(service.ForumService)
	threadService  = new(service.ThreadService)
)
