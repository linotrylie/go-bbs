package requests

type PostSearchRequest struct {
	Fid     int    ` json:"fid"`
	Pid     int    ` json:"pid"`
	Message string ` json:"message"`
}
