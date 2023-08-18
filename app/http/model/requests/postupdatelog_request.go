package requests

type PostUpdateLogRequest struct {
	Logid      int    ` json:"logid"`
	Pid        int    ` json:"pid"`
	Reason     string ` json:"reason"`
	Message    string ` json:"message"`
	CreateDate int    ` json:"createdate"`
	Uid        int    ` json:"uid"`
}
