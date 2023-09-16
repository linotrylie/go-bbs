package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"math"
	"regexp"
)

type UserLogin struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	CaptchaVerify
}

func (param *UserLogin) Validate() error {
	return validation.ValidateStruct(param,
		validation.Field(&param.Username,
			validation.Required.Error("用户名必填！"),
			validation.Length(2, 32).Error("用户名超出规定长度"),
		),
		validation.Field(&param.Password,
			validation.Required.Error("密码必填！"),
			validation.Length(32, 32).Error("密码超出规定长度"),
		),
	)
}

type UserDetail struct {
	Uid int `json:"uid" form:"uid"`
}

func (param *UserDetail) Validate() error {
	return validation.ValidateStruct(param,
		validation.Field(&param.Uid,
			validation.Required.Error("缺少用户id"),
			validation.Min(1).Exclusive().Error("用户id不规范"),
			validation.Max(math.MaxInt).Exclusive().Error("用户id不规范"),
		),
	)
}

type UserChangePassword struct {
	NewPassword       string `json:"new_password,omitempty"`
	OldPassword       string `json:"old_password,omitempty"`
	NewPasswordVerify string `json:"new_password_verify,omitempty"`
}

func (param *UserChangePassword) Validate() error {
	return validation.ValidateStruct(param,
		validation.Field(&param.NewPassword,
			validation.Required.Error("请填写新密码！"),
			validation.Length(32, 32).Error("密码超出规定长度"),
		),
		validation.Field(&param.OldPassword,
			validation.Required.Error("请填写旧密码！"),
			validation.Length(32, 32).Error("密码超出规定长度"),
		),
		validation.Field(&param.NewPasswordVerify,
			//validation.
			validation.Required.Error("请填写重复新密码！"),
			validation.Length(32, 32).Error("密码超出规定长度"),
		),
	)
}

type UserEdit struct {
	Realname string `json:"realname"` // 用户名
	Mobile   string `json:"mobile"`   // 手机号
	Qq       string `json:"qq"`       // QQ
	Email    string `json:"email"`    // 邮箱
	Value    string `json:"value"`    //如果修改邮箱，必须有邮箱验证码
}

func (param *UserEdit) Validate() error {
	return validation.ValidateStruct(param,
		validation.Field(&param.Realname,
			validation.Length(4, 32).Error("超出规定长度！"),
		),
		validation.Field(&param.Mobile,
			validation.Length(11, 11).Error("手机号码超出规定长度！"),
			validation.Match(regexp.MustCompile("^(13[0-9]|14[5|7]|15[0|1|2|3|4|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\\d{8}$")).Error("手机号码格式错误！"),
		),
		validation.Field(&param.Email, validation.Length(8, 64).Error("邮箱超出规定长度！"), is.Email),
		validation.Field(&param.Qq, validation.Match(regexp.MustCompile("[1-9][0-9]{4,}")).Error("QQ号格式不正确！")),
		validation.Field(&param.Value, validation.Length(4, 6).Error("验证码超出规定长度！"), is.UTFDigit),
	)
}

type UserRegister struct {
	Username       string             `json:"username,omitempty"`
	Email          EmailCaptchaVerify `json:"email_verify,omitempty"`
	Captcha        CaptchaVerify      `json:"captcha_verify,omitempty"`
	Password       string             `json:"password,omitempty"`
	PasswordVerify string             `json:"password_verify,omitempty"`
}

func (param *UserRegister) Validate() error {
	return validation.ValidateStruct(param,
		validation.Field(&param.Username,
			validation.Required.Error("用户名必填！"),
			validation.Length(4, 32).Error("用户名超出规定长度"),
		),
		validation.Field(&param.Email),
		validation.Field(&param.Password,
			validation.Required.Error("请填写新密码！"),
			validation.Length(32, 32).Error("密码超出规定长度"),
		),
		validation.Field(&param.PasswordVerify,
			validation.Required.Error("请填写重复新密码！"),
			validation.Length(32, 32).Error("密码超出规定长度"),
		),
		validation.Field(&param.Captcha),
	)
}
