package requests

type XnVoteDetailRequest struct {
	Id       int ` json:"id"`
	VoteId   int ` json:"voteid"`
	Oid      int ` json:"oid"`
	Tid      int ` json:"tid"`
	Uid      int ` json:"uid"`
	VoteTime int ` json:"votetime"`
}
