package requests

type TagCateRequest struct {
	Cateid       int    ` json:"cateid"`
	Fid          int    ` json:"fid"`
	Name         string ` json:"name"`
	Rank         int    ` json:"rank"`
	Enable       int    ` json:"enable"`
	Defaulttagid int    ` json:"defaulttagid"`
	Isforce      int    ` json:"isforce"`
}
