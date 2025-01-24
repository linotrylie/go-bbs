package requests

type KadaoDataRequest struct {
	Kid        int    ` json:"kid"`
	Uid        int    ` json:"uid"`
	Title      string ` json:"title"`
	Dpi        string ` json:"dpi"` // 分辨率
	Data       string ` json:"data"`
	Golds      int    ` json:"golds"` // 加载此方案所需金币数
	IsShare    int    ` json:"isshare"`
	CreateTime string ` json:"createtime"`
	LoadNums   int    ` json:"loadnums"`
	IsShow     int    ` json:"isshow"`
}
