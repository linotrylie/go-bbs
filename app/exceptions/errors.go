package exceptions

import "errors"

var (
	ParamInvalid  = errors.New("参数不正确！")
	DuplicateUser = errors.New("已存在相同用户！")
	UserNotFound  = errors.New("不存在的用户！")
	FailedVerify  = errors.New("验证失败！")
	NotFoundData  = errors.New("未查询到记录！")
	TokenInvalid  = errors.New("token已失效")
)
