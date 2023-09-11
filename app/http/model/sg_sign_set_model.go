package model

import (
	"fmt"
)

type SgSignSet struct {
	changes   map[string]interface{}
	Id        int    `gorm:"primaryKey;column:id" json:"id"`        // id
	SgSignnum int    `gorm:"column:sg_signnum" json:"sg_signnum"`   // 签到总数
	SgSign    int    `gorm:"column:sg_sign" json:"sg_sign"`         // 今日签到人数
	SgSignOne string `gorm:"column:sg_sign_one" json:"sg_sign_one"` // 今日第一
	SgSignTop string `gorm:"column:sg_sign_top" json:"sg_sign_top"` // 今日前十
	Time      int    `gorm:"column:time" json:"time"`               // 最后签到时间
}

func (*SgSignSet) TableName() string {
	return "bbs_sg_sign_set"
}

// Location .
func (obj *SgSignSet) Location() map[string]interface{} {
	return map[string]interface{}{"id": obj.Id}
}

// Redis Key .
func (obj *SgSignSet) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Id)
}

// GetChanges .
func (obj *SgSignSet) GetChanges() map[string]interface{} {
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
func (obj *SgSignSet) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
