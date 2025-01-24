package requests

type ThreadRequest struct {
	Fid            int    ` json:"fid"`
	Tid            int    ` json:"tid"`
	Top            int    ` json:"top"`
	Uid            int    ` json:"uid"`
	Userip         int    ` json:"userip"`
	Subject        string ` json:"subject"`
	CreateDate     int    ` json:"createdate"`
	LastDate       int    ` json:"lastdate"`
	Views          int    ` json:"views"`
	Posts          int    ` json:"posts"`
	Images         int    ` json:"images"`
	Files          int    ` json:"files"`
	Mods           int    ` json:"mods"`
	Closed         int    ` json:"closed"`
	Firstpid       int    ` json:"firstpid"`
	Lastuid        int    ` json:"lastuid"`
	Lastpid        int    ` json:"lastpid"`
	LocationTr     string ` json:"locationtr"`
	Favorites      int    ` json:"favorites"` // 收藏数
	Likes          int    ` json:"likes"`     // 点赞数
	Highlight      int    ` json:"highlight"`
	ContentBuy     int    ` json:"contentbuy"`
	ContentBuyType int    ` json:"contentbuytype"`
	Digest         int    ` json:"digest"`
	Deleted        int    ` json:"deleted"`
	Readp          int    ` json:"readp"`
	OfferNum       int    ` json:"offernum"`
	OfferStatus    int    ` json:"offerstatus"`
	Tagids         string ` json:"tagids"`
	TagidsTime     int    ` json:"tagidstime"`
	IsVote         int    ` json:"isvote"`
	AttachGolds    int    ` json:"attachgolds"`
	ContentGolds   int    ` json:"contentgolds"`
	AdExpireTime   int    ` json:"adexpiretime"`
	IsAd           int    ` json:"isad"`
	AdDays         int    ` json:"addays"`
	IsRaffle       int    ` json:"israffle"`
	RaffleTime     int    ` json:"raffletime"` // 开奖时间
	RaffleUid      string ` json:"raffleuid"`  // 预设中奖用户
	RaffleNums     int    ` json:"rafflenums"` // 中奖用户数
}
