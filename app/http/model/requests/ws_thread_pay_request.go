package requests

type WsThreadPayRequest struct {
	Tid     int ` json:"tid"`     // 帖子id
	Uid     int ` json:"uid"`     // 用户id
	Coin    int ` json:"coin"`    // 支付金币
	Type    int ` json:"type"`    // 支付类型1内容付费2附件付费
	Paytime int ` json:"paytime"` // 支付时间
}
