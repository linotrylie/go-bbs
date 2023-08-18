package requests

type GitTagsRequest struct {
	Tagid int    ` json:"tagid"`
	Name  string ` json:"name"`
	Link  int    ` json:"link"`
}
