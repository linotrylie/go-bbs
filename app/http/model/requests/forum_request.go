package requests

type ForumRequest struct {
	Fid           int    ` json:"fid"`
	Name          string ` json:"name"`
	Rank          int    ` json:"rank"`
	Threads       int    ` json:"threads"`
	Todayposts    int    ` json:"todayposts"`
	Todaythreads  int    ` json:"todaythreads"`
	Brief         string ` json:"brief"`
	Announcement  string ` json:"announcement"`
	Accesson      int    ` json:"accesson"`
	Orderby       int    ` json:"orderby"`
	CreateDate    int    ` json:"createdate"`
	Icon          int    ` json:"icon"`
	Moduids       string ` json:"moduids"`
	SeoTitle      string ` json:"seotitle"`
	SeoKeywords   string ` json:"seokeywords"`
	Digests       int    ` json:"digests"`
	CreateCredits int    ` json:"createcredits"`
	CreateGolds   int    ` json:"creategolds"`
	PostCredits   int    ` json:"postcredits"`
	PostGolds     int    ` json:"postgolds"`
	AllowOffer    int    ` json:"allowoffer"`
}
