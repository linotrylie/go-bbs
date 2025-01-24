package requests

type BuyCardRequest struct {
	Bid         int     ` json:"bid"` // id
	Uid         int     ` json:"uid"` // 用户id
	OrderNo     string  ` json:"orderno"`
	PayType     int     ` json:"paytype"`     // 支付类型 1 微信 2 支付宝
	PayCode     string  ` json:"paycode"`     // 支付码 这是后台审核时要核对的依据
	Type        int     ` json:"type"`        // 充值类型 1 充值卡密 2 充值卡刀软件会员
	Content     string  ` json:"content"`     // 当充值类型是卡刀软件会员时增加的卡刀软件使用时长 1 一个月 2 一个季度 3 一年
	Status      int     ` json:"status"`      // 审核状态 0 未审核 1 已审核 2 系统驳回 3 用户取消
	CreateDate  int     ` json:"createdate"`  // 支付时间
	CheckDate   int     ` json:"checkdate"`   // 审核通过时间
	PayPrice    float64 ` json:"payprice"`    // 支付金额
	PayRatio    float64 ` json:"payratio"`    // 折扣比率
	ContentType int     ` json:"contenttype"` // 充值类型
	Month       int     ` json:"month"`
}
