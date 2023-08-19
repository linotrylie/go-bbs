package requests

type QueueRequest struct {
	Queueid int ` json:"queueid"`
	V       int ` json:"v"`
	Expiry  int ` json:"expiry"`
}
