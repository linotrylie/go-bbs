package requests

type KvRequest struct {
	Name   string ` json:"name"`
	K      string ` json:"k"`
	V      string ` json:"v"`
	Expiry int    ` json:"expiry"`
}
