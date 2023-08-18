package requests

type FriendlinkRequest struct {
	Linkid     int    ` json:"linkid"`
	Type       int    ` json:"type"`
	Rank       int    ` json:"rank"`
	CreateDate int    ` json:"createdate"`
	Name       string ` json:"name"`
	Url        string ` json:"url"`
}
