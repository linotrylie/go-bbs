package requests

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type CaptchaVerify struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (param CaptchaVerify) Validate() error {
	return validation.ValidateStruct(&param,
		validation.Field(&param.Key, validation.Required, validation.Length(16, 64), is.UTFLetterNumeric),
		validation.Field(&param.Value, validation.Required, validation.Length(4, 6), is.UTFDigit),
	)
}
