package requests

type ThreadSearchRequest struct {
	Fid     int    ` json:"fid"`
	Tid     int    ` json:"tid"`
	Message string ` json:"message"`
}
