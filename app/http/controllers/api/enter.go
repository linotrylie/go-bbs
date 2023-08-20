package api

import "go-bbs/app/service"

type ApiGroup struct {
	AuthController
}

var captchaService = new(service.CaptchaService)
