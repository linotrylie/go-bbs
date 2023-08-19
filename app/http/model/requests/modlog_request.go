package requests

type ModlogRequest struct {
	Logid      int    ` json:"logid"`
	Uid        int    ` json:"uid"`
	Tid        int    ` json:"tid"`
	Pid        int    ` json:"pid"`
	Subject    string ` json:"subject"`
	Comment    string ` json:"comment"`
	Rmbs       int    ` json:"rmbs"`
	CreateDate int    ` json:"createdate"`
	Action     string ` json:"action"`
}
