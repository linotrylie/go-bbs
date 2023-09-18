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
	Time      int64  `gorm:"column:time" json:"time"`               // 最后签到时间
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
func (obj *SgSignSet) SetId(val int) *SgSignSet {
	obj.Id += val
	obj.Update("id", obj.Id)
	return obj
}
func (obj *SgSignSet) SetSgSignnum(val int) *SgSignSet {
	obj.SgSignnum += val
	obj.Update("sg_signnum", obj.SgSignnum)
	return obj
}
func (obj *SgSignSet) SetSgSign(val int) *SgSignSet {
	obj.SgSign += val
	obj.Update("sg_sign", obj.SgSign)
	return obj
}
func (obj *SgSignSet) SetSgSignOne(val string) *SgSignSet {
	obj.SgSignOne = val
	obj.Update("sg_sign_one", obj.SgSignOne)
	return obj
}
func (obj *SgSignSet) SetSgSignTop(val string) *SgSignSet {
	obj.SgSignTop = val
	obj.Update("sg_sign_top", obj.SgSignTop)
	return obj
}
func (obj *SgSignSet) SetTime(val int64) *SgSignSet {
	obj.Time += val
	obj.Update("time", obj.Time)
	return obj
}
