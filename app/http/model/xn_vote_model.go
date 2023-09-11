package model

import (
	"fmt"
)

type XnVote struct {
	changes    map[string]interface{}
	VoteId     int    `gorm:"primaryKey;column:vote_id" json:"vote_id"`
	Tid        int    `gorm:"column:tid" json:"tid"`
	Uid        int    `gorm:"column:uid" json:"uid"`
	CreateTime int    `gorm:"column:create_time" json:"create_time"`
	FinishTime int    `gorm:"column:finish_time" json:"finish_time"`
	UpdateTime int    `gorm:"column:update_time" json:"update_time"`
	Type       int    `gorm:"column:type" json:"type"`
	Max        int    `gorm:"column:max" json:"max"`
	Subject    string `gorm:"column:subject" json:"subject"`
}

func (*XnVote) TableName() string {
	return "bbs_xn_vote"
}

// Location .
func (obj *XnVote) Location() map[string]interface{} {
	return map[string]interface{}{"voteid": obj.VoteId}
}

// Redis Key .
func (obj *XnVote) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.VoteId)
}

// GetChanges .
func (obj *XnVote) GetChanges() map[string]interface{} {
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
func (obj *XnVote) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
