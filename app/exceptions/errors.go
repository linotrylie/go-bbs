package exceptions

import "errors"

var (
	ParamInvalid  = errors.New("参数不正确！")
	DuplicateUser = errors.New("已存在相同用户！")
	UserNotFound  = errors.New("不存在的用户！")
	FailedVerify  = errors.New("验证失败！")
	NotFoundData  = errors.New("未查询到记录！")
	TokenInvalid  = errors.New("token已失效")
	ModifyError   = errors.New("修改失败")
	CreateError   = errors.New("新增失败")
	LogBackIn     = errors.New("重新登录")
	NotAuth       = errors.New("没有权限！")
	ThreadIsValid = errors.New("帖子不存在或已被关闭！")
)
