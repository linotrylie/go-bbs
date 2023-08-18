package exceptions

import "errors"

var FailedVerify = errors.New("验证失败！")
var NotFoundData = errors.New("未查询到记录！")
