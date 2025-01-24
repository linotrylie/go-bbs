package requests

type XnVoteInfoRequest struct {
	Oid     int    ` json:"oid"`
	VoteId  int    ` json:"voteid"`
	Tid     int    ` json:"tid"`
	Content string ` json:"content"`
}
