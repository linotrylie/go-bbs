package model

import (
	"fmt"
)

type ThreadTop struct {
	changes map[string]interface{}
	Fid     int `gorm:"column:fid" json:"fid"`
	Tid     int `gorm:"primaryKey;column:tid" json:"tid"`
	Top     int `gorm:"column:top" json:"top"`
}

func (*ThreadTop) TableName() string {
	return "bbs_thread_top"
}

// Location .
func (obj *ThreadTop) Location() map[string]interface{} {
	return map[string]interface{}{"tid": obj.Tid}
}

// Redis Key .
func (obj *ThreadTop) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Tid)
}

// GetChanges .
func (obj *ThreadTop) GetChanges() map[string]interface{} {
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
func (obj *ThreadTop) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
func (obj *ThreadTop) SetFid(val int) *ThreadTop {
	obj.Fid = val
	obj.Update("fid", obj.Fid)
	return obj
}
func (obj *ThreadTop) SetTid(val int) *ThreadTop {
	obj.Tid = val
	obj.Update("tid", obj.Tid)
	return obj
}
func (obj *ThreadTop) SetTop(val int) *ThreadTop {
	obj.Top += val
	obj.Update("top", obj.Top)
	return obj
}
