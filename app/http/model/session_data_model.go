package model

import (
	"fmt"
)

type SessionData struct {
	changes  map[string]interface{}
	Sid      string `gorm:"primaryKey;column:sid" json:"sid"`
	LastDate int64  `gorm:"column:last_date" json:"last_date"`
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
func (obj *SessionData) SetSid(val string) *SessionData {
	obj.Sid = val
	obj.Update("sid", obj.Sid)
	return obj
}
func (obj *SessionData) SetLastDate(val int64) *SessionData {
	obj.LastDate += val
	obj.Update("last_date", obj.LastDate)
	return obj
}
func (obj *SessionData) SetData(val string) *SessionData {
	obj.Data = val
	obj.Update("data", obj.Data)
	return obj
}
