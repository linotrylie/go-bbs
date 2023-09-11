package model

import (
	"fmt"
)

type GitTagsThread struct {
	changes map[string]interface{}
	Tagid   int `gorm:"primaryKey;column:tagid" json:"tagid"`
	Tid     int `gorm:"primaryKey;column:tid" json:"tid"`
}

func (*GitTagsThread) TableName() string {
	return "bbs_git_tags_thread"
}

// Location .
func (obj *GitTagsThread) Location() map[string]interface{} {
	return map[string]interface{}{"tagid": obj.Tagid, "tid": obj.Tid}
}

// Redis Key .
func (obj *GitTagsThread) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Tagid) + "_" + fmt.Sprintf("%v", obj.Tid)
}

// GetChanges .
func (obj *GitTagsThread) GetChanges() map[string]interface{} {
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
func (obj *GitTagsThread) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
