package model

import (
	"fmt"
)

type Mythread struct {
	changes map[string]interface{}
	Uid     int `gorm:"primaryKey;column:uid" json:"uid"`
	Tid     int `gorm:"primaryKey;column:tid" json:"tid"`
}

func (*Mythread) TableName() string {
	return "bbs_mythread"
}

// Location .
func (obj *Mythread) Location() map[string]interface{} {
	return map[string]interface{}{"uid": obj.Uid, "tid": obj.Tid}
}

// Redis Key .
func (obj *Mythread) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Tid) + "_" + fmt.Sprintf("%v", obj.Uid)
}

// GetChanges .
func (obj *Mythread) GetChanges() map[string]interface{} {
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
func (obj *Mythread) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
func (obj *Mythread) SetUid(val int) *Mythread {
	obj.Uid = val
	obj.Update("uid", obj.Uid)
	return obj
}
func (obj *Mythread) SetTid(val int) *Mythread {
	obj.Tid = val
	obj.Update("tid", obj.Tid)
	return obj
}
