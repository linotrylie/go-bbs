package model

import (
	"fmt"
)

type Mypost struct {
	changes map[string]interface{}
	Uid     int `gorm:"primaryKey;column:uid" json:"uid"`
	Tid     int `gorm:"column:tid" json:"tid"`
	Pid     int `gorm:"primaryKey;column:pid" json:"pid"`
}

func (*Mypost) TableName() string {
	return "bbs_mypost"
}

// Location .
func (obj *Mypost) Location() map[string]interface{} {
	return map[string]interface{}{"uid": obj.Uid, "pid": obj.Pid}
}

// Redis Key .
func (obj *Mypost) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Uid) + "_" + fmt.Sprintf("%v", obj.Pid)
}

// GetChanges .
func (obj *Mypost) GetChanges() map[string]interface{} {
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
func (obj *Mypost) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
func (obj *Mypost) SetUid(val int) *Mypost {
	obj.Uid = val
	obj.Update("uid", obj.Uid)
	return obj
}
func (obj *Mypost) SetTid(val int) *Mypost {
	obj.Tid = val
	obj.Update("tid", obj.Tid)
	return obj
}
func (obj *Mypost) SetPid(val int) *Mypost {
	obj.Pid = val
	obj.Update("pid", obj.Pid)
	return obj
}
