package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"math"
)

type UserLogin struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	CaptchaVerify
}

func (param *UserLogin) Validate() error {
	return validation.ValidateStruct(param,
		validation.Field(&param.Username, validation.Required, validation.Length(4, 32)),
		validation.Field(&param.Password, validation.Required, validation.Length(4, 32)),
	)
}

type UserDetail struct {
	Uid int `json:"uid" form:"uid"`
}

func (param *UserDetail) Validate() error {
	return validation.ValidateStruct(param,
		validation.Field(&param.Uid, validation.Required, validation.Min(1), validation.Max(math.MaxInt)),
	)
}
