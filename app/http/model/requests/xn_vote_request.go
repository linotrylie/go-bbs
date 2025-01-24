package requests

type XnVoteRequest struct {
	VoteId     int    ` json:"voteid"`
	Tid        int    ` json:"tid"`
	Uid        int    ` json:"uid"`
	CreateTime int    ` json:"createtime"`
	FinishTime int    ` json:"finishtime"`
	UpdateTime int    ` json:"updatetime"`
	Type       int    ` json:"type"`
	Max        int    ` json:"max"`
	Subject    string ` json:"subject"`
}
