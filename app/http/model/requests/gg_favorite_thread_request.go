package requests

type GgFavoriteThreadRequest struct {
	Favid int ` json:"favid"`
	Tid   int ` json:"tid"`
	Uid   int ` json:"uid"`
}
