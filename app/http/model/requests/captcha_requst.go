package requests

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type CaptchaVerify struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (param *CaptchaVerify) Validate() error {
	return validation.ValidateStruct(param,
		validation.Field(&param.Key, validation.Required, validation.Length(16, 64), is.UTFLetterNumeric),
		validation.Field(&param.Value, validation.Required, validation.Length(4, 6), is.UTFDigit),
	)
}

type EmailCaptcha struct {
	Email string `json:"email"`
}

func (param *EmailCaptcha) Validate() error {
	return validation.ValidateStruct(param,
		validation.Field(&param.Email, validation.Required, validation.Length(8, 64), is.Email),
	)
}

type EmailCaptchaVerify struct {
	Value string `json:"value"`
	Email string `json:"email"`
}

func (param *EmailCaptchaVerify) Validate() error {
	return validation.ValidateStruct(param,
		validation.Field(&param.Value, validation.Required, validation.Length(4, 6), is.UTFDigit),
		validation.Field(&param.Email, validation.Required, validation.Length(8, 64), is.Email),
	)
}
