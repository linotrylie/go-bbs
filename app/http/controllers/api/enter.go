package api

import "GoFreeBns/app/service"

type ApiGroup struct {
	AuthController
}

var captchaService = new(service.CaptchaService)
