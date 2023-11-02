package model

import (
	"fmt"
	"time"
)

type KadaoData struct {
	changes    map[string]interface{}
	Kid        int       `gorm:"primaryKey;column:kid" json:"kid"`
	Uid        int       `gorm:"column:uid" json:"uid"`
	Title      string    `gorm:"column:title" json:"title"`
	Data       string    `gorm:"column:data" json:"data"`
	IsShare    int       `gorm:"column:is_share" json:"is_share"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
}

func (*KadaoData) TableName() string {
	return "bbs_kadao_data"
}

// Location .
func (obj *KadaoData) Location() map[string]interface{} {
	return map[string]interface{}{"kid": obj.Kid}
}

// Redis Key .
func (obj *KadaoData) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Kid)
}

// GetChanges .
func (obj *KadaoData) GetChanges() map[string]interface{} {
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
func (obj *KadaoData) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
func (obj *KadaoData) SetKid(val int) *KadaoData {
	obj.Kid = val
	obj.Update("kid", obj.Kid)
	return obj
}
func (obj *KadaoData) SetUid(val int) *KadaoData {
	obj.Uid = val
	obj.Update("uid", obj.Uid)
	return obj
}
func (obj *KadaoData) SetTitle(val string) *KadaoData {
	obj.Title = val
	obj.Update("title", obj.Title)
	return obj
}
func (obj *KadaoData) SetData(val string) *KadaoData {
	obj.Data = val
	obj.Update("data", obj.Data)
	return obj
}
func (obj *KadaoData) SetIsShare(val int) *KadaoData {
	obj.IsShare += val
	obj.Update("is_share", obj.IsShare)
	return obj
}
func (obj *KadaoData) SetCreateTime(val time.Time) *KadaoData {
	obj.CreateTime = val
	obj.Update("create_time", obj.CreateTime)
	return obj
}
