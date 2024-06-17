package model

import (
	"fmt"
)

type KadaoData struct {
	changes    map[string]interface{}
	Kid        int    `gorm:"primaryKey;column:kid" json:"kid"`
	Uid        int    `gorm:"column:uid" json:"uid"`
	Title      string `gorm:"column:title" json:"title"`
	Dpi        string `gorm:"column:dpi" json:"dpi"` // 分辨率
	Data       string `gorm:"column:data" json:"data"`
	Golds      int    `gorm:"column:golds" json:"golds"` // 加载此方案所需金币数
	IsShare    int    `gorm:"column:is_share" json:"is_share"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
	LoadNums   int    `gorm:"column:load_nums" json:"load_nums"`
	IsShow     int    `gorm:"column:is_show" json:"is_show"`
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
func (obj *KadaoData) SetDpi(val string) *KadaoData {
	obj.Dpi = val
	obj.Update("dpi", obj.Dpi)
	return obj
}
func (obj *KadaoData) SetData(val string) *KadaoData {
	obj.Data = val
	obj.Update("data", obj.Data)
	return obj
}
func (obj *KadaoData) SetGolds(val int) *KadaoData {
	obj.Golds += val
	obj.Update("golds", obj.Golds)
	return obj
}
func (obj *KadaoData) SetIsShare(val int) *KadaoData {
	obj.IsShare += val
	obj.Update("is_share", obj.IsShare)
	return obj
}
func (obj *KadaoData) SetCreateTime(val int64) *KadaoData {
	obj.CreateTime += val
	obj.Update("create_time", obj.CreateTime)
	return obj
}
func (obj *KadaoData) SetLoadNums(val int) *KadaoData {
	obj.LoadNums += val
	obj.Update("load_nums", obj.LoadNums)
	return obj
}
func (obj *KadaoData) SetIsShow(val int) *KadaoData {
	obj.IsShow += val
	obj.Update("is_show", obj.IsShow)
	return obj
}
