package model

import (
	"fmt"
)

type Kv struct {
	changes map[string]interface{}
	K       string `gorm:"primaryKey;column:k" json:"k"`
	V       string `gorm:"column:v" json:"v"`
	Expiry  int    `gorm:"column:expiry" json:"expiry"`
}

func (*Kv) TableName() string {
	return "bbs_kv"
}

// Location .
func (obj *Kv) Location() map[string]interface{} {
	return map[string]interface{}{"k": obj.K}
}

// Redis Key .
func (obj *Kv) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.K)
}

// GetChanges .
func (obj *Kv) GetChanges() map[string]interface{} {
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
func (obj *Kv) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
