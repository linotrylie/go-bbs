package model

import (
	"fmt"
)

type XnVoteDetail struct {
	changes  map[string]interface{}
	Id       int   `gorm:"primaryKey;column:id" json:"id"`
	VoteId   int   `gorm:"column:vote_id" json:"vote_id"`
	Oid      int   `gorm:"column:oid" json:"oid"`
	Tid      int   `gorm:"column:tid" json:"tid"`
	Uid      int   `gorm:"column:uid" json:"uid"`
	VoteTime int64 `gorm:"column:vote_time" json:"vote_time"`
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
func (obj *XnVoteDetail) SetId(val int) *XnVoteDetail {
	obj.Id += val
	obj.Update("id", obj.Id)
	return obj
}
func (obj *XnVoteDetail) SetVoteId(val int) *XnVoteDetail {
	obj.VoteId += val
	obj.Update("vote_id", obj.VoteId)
	return obj
}
func (obj *XnVoteDetail) SetOid(val int) *XnVoteDetail {
	obj.Oid = val
	obj.Update("oid", obj.Oid)
	return obj
}
func (obj *XnVoteDetail) SetTid(val int) *XnVoteDetail {
	obj.Tid = val
	obj.Update("tid", obj.Tid)
	return obj
}
func (obj *XnVoteDetail) SetUid(val int) *XnVoteDetail {
	obj.Uid = val
	obj.Update("uid", obj.Uid)
	return obj
}
func (obj *XnVoteDetail) SetVoteTime(val int64) *XnVoteDetail {
	obj.VoteTime += val
	obj.Update("vote_time", obj.VoteTime)
	return obj
}
