package requests

type PaylistRequest struct {
	Plid       int ` json:"plid"`
	Tid        int ` json:"tid"`        // tid
	Uid        int ` json:"uid"`        // uid
	Num        int ` json:"num"`        // pay_anycredit_num
	CreditType int ` json:"credittype"` // 1exp_2gold_3rmb
	Type       int ` json:"type"`
	Rate       int ` json:"rate"`
	Paytime    int ` json:"paytime"` // time
}
