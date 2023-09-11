package model

import (
	"fmt"
)

type XnVoteInfo struct {
	changes map[string]interface{}
	Oid     int    `gorm:"primaryKey;column:oid" json:"oid"`
	VoteId  int    `gorm:"column:vote_id" json:"vote_id"`
	Tid     int    `gorm:"column:tid" json:"tid"`
	Content string `gorm:"column:content" json:"content"`
}

func (*XnVoteInfo) TableName() string {
	return "bbs_xn_vote_info"
}

// Location .
func (obj *XnVoteInfo) Location() map[string]interface{} {
	return map[string]interface{}{"oid": obj.Oid}
}

// Redis Key .
func (obj *XnVoteInfo) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Oid)
}

// GetChanges .
func (obj *XnVoteInfo) GetChanges() map[string]interface{} {
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
func (obj *XnVoteInfo) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
