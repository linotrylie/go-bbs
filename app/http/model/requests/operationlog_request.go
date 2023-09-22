package requests

type OperationLogRequest struct {
	Id           int    ` json:"id"`
	Ip           int    ` json:"ip"`
	Method       string ` json:"method"`  // 请求方法
	Path         string ` json:"path"`    // 请求路径
	Status       int    ` json:"status"`  // 请求状态
	Latency      int    ` json:"latency"` // 延迟
	Agent        string ` json:"agent"`
	ErrorMessage string ` json:"errormessage"`
	Body         string ` json:"body"`
	Resp         string ` json:"resp"`
	Uid          int    ` json:"uid"`
}
