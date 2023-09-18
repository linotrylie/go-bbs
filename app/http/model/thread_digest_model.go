package model

import (
	"fmt"
)

type ThreadDigest struct {
	changes map[string]interface{}
	Fid     int `gorm:"column:fid" json:"fid"`
	Tid     int `gorm:"primaryKey;column:tid" json:"tid"`
	Uid     int `gorm:"column:uid" json:"uid"`
	Digest  int `gorm:"column:digest" json:"digest"`
}

func (*ThreadDigest) TableName() string {
	return "bbs_thread_digest"
}

// Location .
func (obj *ThreadDigest) Location() map[string]interface{} {
	return map[string]interface{}{"tid": obj.Tid}
}

// Redis Key .
func (obj *ThreadDigest) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Tid)
}

// GetChanges .
func (obj *ThreadDigest) GetChanges() map[string]interface{} {
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
func (obj *ThreadDigest) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
func (obj *ThreadDigest) SetFid(val int) *ThreadDigest {
	obj.Fid = val
	obj.Update("fid", obj.Fid)
	return obj
}
func (obj *ThreadDigest) SetTid(val int) *ThreadDigest {
	obj.Tid = val
	obj.Update("tid", obj.Tid)
	return obj
}
func (obj *ThreadDigest) SetUid(val int) *ThreadDigest {
	obj.Uid = val
	obj.Update("uid", obj.Uid)
	return obj
}
func (obj *ThreadDigest) SetDigest(val int) *ThreadDigest {
	obj.Digest += val
	obj.Update("digest", obj.Digest)
	return obj
}
