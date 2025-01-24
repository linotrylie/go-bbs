package requests

type KadaoDataShareLogRequest struct {
	Id         int ` json:"id"`
	Kid        int ` json:"kid"`     // 卡刀数据id
	FromUid    int ` json:"fromuid"` // 分享者
	ToUid      int ` json:"touid"`   // 购买者
	Golds      int ` json:"golds"`   // 花费金币
	CreateTime int ` json:"createtime"`
}
