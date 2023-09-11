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
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Queueid) + "_" + fmt.Sprintf("%v", obj.V)
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
