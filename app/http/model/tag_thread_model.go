package model

import (
	"fmt"
)

type TagThread struct {
	changes map[string]interface{}
	Tagid   int `gorm:"primaryKey;column:tagid" json:"tagid"`
	Tid     int `gorm:"primaryKey;column:tid" json:"tid"`
}

func (*TagThread) TableName() string {
	return "bbs_tag_thread"
}

// Location .
func (obj *TagThread) Location() map[string]interface{} {
	return map[string]interface{}{"tagid": obj.Tagid, "tid": obj.Tid}
}

// Redis Key .
func (obj *TagThread) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Tid) + "_" + fmt.Sprintf("%v", obj.Tagid)
}

// GetChanges .
func (obj *TagThread) GetChanges() map[string]interface{} {
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
func (obj *TagThread) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
func (obj *TagThread) SetTagid(val int) *TagThread {
	obj.Tagid = val
	obj.Update("tagid", obj.Tagid)
	return obj
}
func (obj *TagThread) SetTid(val int) *TagThread {
	obj.Tid = val
	obj.Update("tid", obj.Tid)
	return obj
}
