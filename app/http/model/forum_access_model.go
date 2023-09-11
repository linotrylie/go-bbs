package model

import (
	"fmt"
)

type ForumAccess struct {
	changes     map[string]interface{}
	Fid         int `gorm:"primaryKey;column:fid" json:"fid"`
	Gid         int `gorm:"primaryKey;column:gid" json:"gid"`
	Allowread   int `gorm:"column:allowread" json:"allowread"`
	Allowthread int `gorm:"column:allowthread" json:"allowthread"`
	Allowpost   int `gorm:"column:allowpost" json:"allowpost"`
	Allowattach int `gorm:"column:allowattach" json:"allowattach"`
	Allowdown   int `gorm:"column:allowdown" json:"allowdown"`
}

func (*ForumAccess) TableName() string {
	return "bbs_forum_access"
}

// Location .
func (obj *ForumAccess) Location() map[string]interface{} {
	return map[string]interface{}{"fid": obj.Fid, "gid": obj.Gid}
}

// Redis Key .
func (obj *ForumAccess) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Fid) + "_" + fmt.Sprintf("%v", obj.Gid)
}

// GetChanges .
func (obj *ForumAccess) GetChanges() map[string]interface{} {
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
func (obj *ForumAccess) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
