package model

import (
	"fmt"
)

type UserPay struct {
	changes    map[string]interface{}
	Cid        int    `gorm:"primaryKey;column:cid" json:"cid"`
	Uid        int    `gorm:"column:uid" json:"uid"`
	Status     int    `gorm:"column:status" json:"status"`
	Num        int    `gorm:"column:num" json:"num"`
	Type       int    `gorm:"column:type" json:"type"`
	CreditType int    `gorm:"column:credit_type" json:"credit_type"`
	Code       string `gorm:"column:code" json:"code"`
	Time       int64  `gorm:"column:time" json:"time"`
}

func (*UserPay) TableName() string {
	return "bbs_user_pay"
}

// Location .
func (obj *UserPay) Location() map[string]interface{} {
	return map[string]interface{}{"cid": obj.Cid}
}

// Redis Key .
func (obj *UserPay) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Cid)
}

// GetChanges .
func (obj *UserPay) GetChanges() map[string]interface{} {
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
func (obj *UserPay) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
func (obj *UserPay) SetCid(val int) *UserPay {
	obj.Cid = val
	obj.Update("cid", obj.Cid)
	return obj
}
func (obj *UserPay) SetUid(val int) *UserPay {
	obj.Uid = val
	obj.Update("uid", obj.Uid)
	return obj
}
func (obj *UserPay) SetStatus(val int) *UserPay {
	obj.Status += val
	obj.Update("status", obj.Status)
	return obj
}
func (obj *UserPay) SetNum(val int) *UserPay {
	obj.Num += val
	obj.Update("num", obj.Num)
	return obj
}
func (obj *UserPay) SetType(val int) *UserPay {
	obj.Type += val
	obj.Update("type", obj.Type)
	return obj
}
func (obj *UserPay) SetCreditType(val int) *UserPay {
	obj.CreditType += val
	obj.Update("credit_type", obj.CreditType)
	return obj
}
func (obj *UserPay) SetCode(val string) *UserPay {
	obj.Code = val
	obj.Update("code", obj.Code)
	return obj
}
func (obj *UserPay) SetTime(val int64) *UserPay {
	obj.Time += val
	obj.Update("time", obj.Time)
	return obj
}
