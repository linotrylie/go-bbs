package requests

type SessionDataRequest struct {
	Sid      string ` json:"sid"`
	LastDate int    ` json:"lastdate"`
	Data     string ` json:"data"`
}
