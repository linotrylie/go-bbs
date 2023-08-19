package requests

type InviteRequest struct {
	Uid     int ` json:"uid"`
	Ip      int ` json:"ip"`
	Regtime int ` json:"regtime"`
}
