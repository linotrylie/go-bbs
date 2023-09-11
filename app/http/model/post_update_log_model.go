package model

import (
	"fmt"
)

type PostUpdateLog struct {
	changes    map[string]interface{}
	Logid      int    `gorm:"primaryKey;column:logid" json:"logid"`
	Pid        int    `gorm:"column:pid" json:"pid"`
	Reason     string `gorm:"column:reason" json:"reason"`
	Message    string `gorm:"column:message" json:"message"`
	CreateDate int    `gorm:"column:create_date" json:"create_date"`
	Uid        int    `gorm:"column:uid" json:"uid"`
}

func (*PostUpdateLog) TableName() string {
	return "bbs_post_update_log"
}

// Location .
func (obj *PostUpdateLog) Location() map[string]interface{} {
	return map[string]interface{}{"logid": obj.Logid}
}

// Redis Key .
func (obj *PostUpdateLog) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Logid)
}

// GetChanges .
func (obj *PostUpdateLog) GetChanges() map[string]interface{} {
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
func (obj *PostUpdateLog) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
