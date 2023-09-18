package model

import (
	"fmt"
)

type Queue struct {
	changes map[string]interface{}
	Queueid int `gorm:"primaryKey;column:queueid" json:"queueid"`
	V       int `gorm:"primaryKey;column:v" json:"v"`
	Expiry  int `gorm:"column:expiry" json:"expiry"`
}

func (*Queue) TableName() string {
	return "bbs_queue"
}

// Location .
func (obj *Queue) Location() map[string]interface{} {
	return map[string]interface{}{"queueid": obj.Queueid, "v": obj.V}
}

// Redis Key .
func (obj *Queue) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.V) + "_" + fmt.Sprintf("%v", obj.Queueid)
}

// GetChanges .
func (obj *Queue) GetChanges() map[string]interface{} {
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
func (obj *Queue) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
func (obj *Queue) SetQueueid(val int) *Queue {
	obj.Queueid = val
	obj.Update("queueid", obj.Queueid)
	return obj
}
func (obj *Queue) SetV(val int) *Queue {
	obj.V += val
	obj.Update("v", obj.V)
	return obj
}
func (obj *Queue) SetExpiry(val int) *Queue {
	obj.Expiry += val
	obj.Update("expiry", obj.Expiry)
	return obj
}
