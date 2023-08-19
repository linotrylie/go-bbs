package requests

type KvRequest struct {
	K      string ` json:"k"`
	V      string ` json:"v"`
	Expiry int    ` json:"expiry"`
}
