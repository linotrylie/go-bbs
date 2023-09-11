package model

import (
	"fmt"
)

type Notice struct {
	changes    map[string]interface{}
	Nid        int    `gorm:"primaryKey;column:nid" json:"nid"`
	Fromuid    int    `gorm:"column:fromuid" json:"fromuid"`
	Recvuid    int    `gorm:"column:recvuid" json:"recvuid"`
	CreateDate int    `gorm:"column:create_date" json:"create_date"`
	Isread     int    `gorm:"column:isread" json:"isread"`
	Type       int    `gorm:"column:type" json:"type"`
	Message    string `gorm:"column:message" json:"message"`
}

func (*Notice) TableName() string {
	return "bbs_notice"
}

// Location .
func (obj *Notice) Location() map[string]interface{} {
	return map[string]interface{}{"nid": obj.Nid}
}

// Redis Key .
func (obj *Notice) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Nid)
}

// GetChanges .
func (obj *Notice) GetChanges() map[string]interface{} {
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
func (obj *Notice) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
