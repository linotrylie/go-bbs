package response

type ThreadVo struct {
	Fid            int     `json:"fid"`
	Tid            int     `json:"tid"`
	Top            int     `json:"top"`
	Uid            int     `json:"uid"`
	Userip         string  `json:"userip"`
	Subject        string  `json:"subject"`
	CreateDate     string  `json:"create_date"`
	LastDate       string  `json:"last_date"`
	Views          int     `json:"views"`
	Posts          int     `json:"posts"`
	Images         int     `json:"images"`
	Files          int     `json:"files"`
	Mods           int     `json:"mods"`
	Closed         int     `json:"closed"`
	Firstpid       int     `json:"firstpid"`
	Lastuid        int     `json:"lastuid"`
	Lastpid        int     `json:"lastpid"`
	LocationTr     string  `json:"location_tr"`
	Favorites      int     `json:"favorites"` // 收藏数
	Likes          int     `json:"likes"`     // 点赞数
	Highlight      int     `json:"highlight"`
	ContentBuy     int     `json:"content_buy"`
	ContentBuyType int     `json:"content_buy_type"`
	Digest         int     `json:"digest"`
	Deleted        int     `json:"deleted"`
	Readp          int     `json:"readp"`
	OfferNum       int     `json:"offernum"`
	OfferStatus    int     `json:"offerstatus"`
	Tagids         string  `json:"tagids"`
	TagidsTime     string  `json:"tagids_time"`
	IsVote         int     `json:"is_vote"`
	ActivityId     int     `json:"activity_id"`
	AttachGolds    int     `json:"attach_golds"`
	ContentGolds   int     `json:"content_golds"`
	User           *UserVo `json:"user,omitempty"`
}
