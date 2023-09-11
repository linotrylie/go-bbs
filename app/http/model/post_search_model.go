package model

import (
	"fmt"
)

type PostSearch struct {
	changes map[string]interface{}
	Fid     int    `gorm:"column:fid" json:"fid"`
	Pid     int    `gorm:"primaryKey;column:pid" json:"pid"`
	Message string `gorm:"column:message" json:"message"`
}

func (*PostSearch) TableName() string {
	return "bbs_post_search"
}

// Location .
func (obj *PostSearch) Location() map[string]interface{} {
	return map[string]interface{}{"pid": obj.Pid}
}

// Redis Key .
func (obj *PostSearch) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Pid)
}

// GetChanges .
func (obj *PostSearch) GetChanges() map[string]interface{} {
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
func (obj *PostSearch) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
