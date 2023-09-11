package model

import (
	"fmt"
)

type Paylist struct {
	changes    map[string]interface{}
	Plid       int `gorm:"primaryKey;column:plid" json:"plid"`
	Tid        int `gorm:"column:tid" json:"tid"`                 // tid
	Uid        int `gorm:"column:uid" json:"uid"`                 // uid
	Num        int `gorm:"column:num" json:"num"`                 // pay_anycredit_num
	CreditType int `gorm:"column:credit_type" json:"credit_type"` // 1exp_2gold_3rmb
	Type       int `gorm:"column:type" json:"type"`
	Rate       int `gorm:"column:rate" json:"rate"`
	Paytime    int `gorm:"column:paytime" json:"paytime"` // time
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
