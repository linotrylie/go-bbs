package requests

type CacheRequest struct {
	K      string ` json:"k"`
	V      string ` json:"v"`
	Expiry int    ` json:"expiry"`
}
