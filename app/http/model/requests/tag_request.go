package requests

type TagRequest struct {
	Tagid  int    ` json:"tagid"`
	Cateid int    ` json:"cateid"`
	Name   string ` json:"name"`
	Rank   int    ` json:"rank"`
	Enable int    ` json:"enable"`
	Style  string ` json:"style"`
}
