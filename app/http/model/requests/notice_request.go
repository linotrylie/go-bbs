package requests

type NoticeRequest struct {
	Nid        int    ` json:"nid"`
	Fromuid    int    ` json:"fromuid"`
	Recvuid    int    ` json:"recvuid"`
	CreateDate int    ` json:"createdate"`
	Isread     int    ` json:"isread"`
	Type       int    ` json:"type"`
	Message    string ` json:"message"`
}
