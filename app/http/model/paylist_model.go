package model

import (
	"fmt"
)

type Paylist struct {
	changes    map[string]interface{}
	Plid       int   `gorm:"primaryKey;column:plid" json:"plid"`
	Tid        int   `gorm:"column:tid" json:"tid"`                 // tid
	Uid        int   `gorm:"column:uid" json:"uid"`                 // uid
	Num        int   `gorm:"column:num" json:"num"`                 // pay_anycredit_num
	CreditType int   `gorm:"column:credit_type" json:"credit_type"` // 1exp_2gold_3rmb
	Type       int   `gorm:"column:type" json:"type"`
	Rate       int   `gorm:"column:rate" json:"rate"`
	Paytime    int64 `gorm:"column:paytime" json:"paytime"` // time
}

func (*Paylist) TableName() string {
	return "bbs_paylist"
}

// Location .
func (obj *Paylist) Location() map[string]interface{} {
	return map[string]interface{}{"plid": obj.Plid}
}

// Redis Key .
func (obj *Paylist) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Plid)
}

// GetChanges .
func (obj *Paylist) GetChanges() map[string]interface{} {
	if obj.changes == nil {
		return nil
	}
	result := make(map[string]interface{})
	for k, v := range obj.changes {
		result[k] = v
	}
	obj.changes = nil
	return result
}

// Update .
func (obj *Paylist) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
func (obj *Paylist) SetPlid(val int) *Paylist {
	obj.Plid = val
	obj.Update("plid", obj.Plid)
	return obj
}
func (obj *Paylist) SetTid(val int) *Paylist {
	obj.Tid = val
	obj.Update("tid", obj.Tid)
	return obj
}
func (obj *Paylist) SetUid(val int) *Paylist {
	obj.Uid = val
	obj.Update("uid", obj.Uid)
	return obj
}
func (obj *Paylist) SetNum(val int) *Paylist {
	obj.Num += val
	obj.Update("num", obj.Num)
	return obj
}
func (obj *Paylist) SetCreditType(val int) *Paylist {
	obj.CreditType += val
	obj.Update("credit_type", obj.CreditType)
	return obj
}
func (obj *Paylist) SetType(val int) *Paylist {
	obj.Type += val
	obj.Update("type", obj.Type)
	return obj
}
func (obj *Paylist) SetRate(val int) *Paylist {
	obj.Rate += val
	obj.Update("rate", obj.Rate)
	return obj
}
func (obj *Paylist) SetPaytime(val int64) *Paylist {
	obj.Paytime += val
	obj.Update("paytime", obj.Paytime)
	return obj
}
