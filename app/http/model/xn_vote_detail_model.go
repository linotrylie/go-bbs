package model

import (
	"fmt"
)

type XnVoteDetail struct {
	changes  map[string]interface{}
	Id       int `gorm:"primaryKey;column:id" json:"id"`
	VoteId   int `gorm:"column:vote_id" json:"vote_id"`
	Oid      int `gorm:"column:oid" json:"oid"`
	Tid      int `gorm:"column:tid" json:"tid"`
	Uid      int `gorm:"column:uid" json:"uid"`
	VoteTime int `gorm:"column:vote_time" json:"vote_time"`
}

func (*XnVoteDetail) TableName() string {
	return "bbs_xn_vote_detail"
}

// Location .
func (obj *XnVoteDetail) Location() map[string]interface{} {
	return map[string]interface{}{"id": obj.Id}
}

// Redis Key .
func (obj *XnVoteDetail) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Id)
}

// GetChanges .
func (obj *XnVoteDetail) GetChanges() map[string]interface{} {
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
func (obj *XnVoteDetail) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
