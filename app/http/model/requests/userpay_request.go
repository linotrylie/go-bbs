package requests

type UserPayRequest struct {
	Cid        int    ` json:"cid"`
	Uid        int    ` json:"uid"`
	Status     int    ` json:"status"`
	Num        int    ` json:"num"`
	Type       int    ` json:"type"`
	CreditType int    ` json:"credittype"`
	Code       string ` json:"code"`
	Time       int    ` json:"time"`
}
