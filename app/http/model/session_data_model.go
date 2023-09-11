package model

import (
	"fmt"
)

type SessionData struct {
	changes  map[string]interface{}
	Sid      string `gorm:"primaryKey;column:sid" json:"sid"`
	LastDate int    `gorm:"column:last_date" json:"last_date"`
	Data     string `gorm:"column:data" json:"data"`
}

func (*SessionData) TableName() string {
	return "bbs_session_data"
}

// Location .
func (obj *SessionData) Location() map[string]interface{} {
	return map[string]interface{}{"sid": obj.Sid}
}

// Redis Key .
func (obj *SessionData) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Sid)
}

// GetChanges .
func (obj *SessionData) GetChanges() map[string]interface{} {
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
func (obj *SessionData) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
