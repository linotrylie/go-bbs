package requests

type OperationLogRequest struct {
	Id           int    ` json:"id"`
	Uid          int    ` json:"uid"`
	Ip           int    ` json:"ip"`
	Method       string ` json:"method"`  // 请求方法
	Path         string ` json:"path"`    // 请求路径
	Status       int    ` json:"status"`  // 请求状态
	Latency      string ` json:"latency"` // 延迟
	Agent        string ` json:"agent"`
	ErrorMessage string ` json:"errormessage"`
	Body         string ` json:"body"`
	Resp         string ` json:"resp"`
}
