package model

import (
	"fmt"
)

type XnVote struct {
	changes    map[string]interface{}
	VoteId     int    `gorm:"primaryKey;column:vote_id" json:"vote_id"`
	Tid        int    `gorm:"column:tid" json:"tid"`
	Uid        int    `gorm:"column:uid" json:"uid"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
	FinishTime int64  `gorm:"column:finish_time" json:"finish_time"`
	UpdateTime int64  `gorm:"column:update_time" json:"update_time"`
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
func (obj *XnVote) SetVoteId(val int) *XnVote {
	obj.VoteId += val
	obj.Update("vote_id", obj.VoteId)
	return obj
}
func (obj *XnVote) SetTid(val int) *XnVote {
	obj.Tid = val
	obj.Update("tid", obj.Tid)
	return obj
}
func (obj *XnVote) SetUid(val int) *XnVote {
	obj.Uid = val
	obj.Update("uid", obj.Uid)
	return obj
}
func (obj *XnVote) SetCreateTime(val int64) *XnVote {
	obj.CreateTime += val
	obj.Update("create_time", obj.CreateTime)
	return obj
}
func (obj *XnVote) SetFinishTime(val int64) *XnVote {
	obj.FinishTime += val
	obj.Update("finish_time", obj.FinishTime)
	return obj
}
func (obj *XnVote) SetUpdateTime(val int64) *XnVote {
	obj.UpdateTime += val
	obj.Update("update_time", obj.UpdateTime)
	return obj
}
func (obj *XnVote) SetType(val int) *XnVote {
	obj.Type += val
	obj.Update("type", obj.Type)
	return obj
}
func (obj *XnVote) SetMax(val int) *XnVote {
	obj.Max += val
	obj.Update("max", obj.Max)
	return obj
}
func (obj *XnVote) SetSubject(val string) *XnVote {
	obj.Subject = val
	obj.Update("subject", obj.Subject)
	return obj
}
