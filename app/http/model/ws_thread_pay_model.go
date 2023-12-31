package model

import (
	"fmt"
	"time"
)

type WsThreadPay struct {
	changes map[string]interface{}
	Tid     int   `gorm:"column:tid" json:"tid"`         // 帖子id
	Uid     int   `gorm:"column:uid" json:"uid"`         // 用户id
	Coin    int   `gorm:"column:coin" json:"coin"`       // 支付金币
	Type    int   `gorm:"column:type" json:"type"`       // 支付类型1内容付费2附件付费
	Paytime int64 `gorm:"column:paytime" json:"paytime"` // 支付时间
}

func (*WsThreadPay) TableName() string {
	return "bbs_ws_thread_pay"
}

// Location .
func (obj *WsThreadPay) Location() map[string]interface{} {
	return map[string]interface{}{}
}

// Redis Key .
func (obj *WsThreadPay) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", time.Now().Unix())
}

// GetChanges .
func (obj *WsThreadPay) GetChanges() map[string]interface{} {
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
func (obj *WsThreadPay) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
func (obj *WsThreadPay) SetTid(val int) *WsThreadPay {
	obj.Tid = val
	obj.Update("tid", obj.Tid)
	return obj
}
func (obj *WsThreadPay) SetUid(val int) *WsThreadPay {
	obj.Uid = val
	obj.Update("uid", obj.Uid)
	return obj
}
func (obj *WsThreadPay) SetCoin(val int) *WsThreadPay {
	obj.Coin += val
	obj.Update("coin", obj.Coin)
	return obj
}
func (obj *WsThreadPay) SetType(val int) *WsThreadPay {
	obj.Type += val
	obj.Update("type", obj.Type)
	return obj
}
func (obj *WsThreadPay) SetPaytime(val int64) *WsThreadPay {
	obj.Paytime += val
	obj.Update("paytime", obj.Paytime)
	return obj
}
