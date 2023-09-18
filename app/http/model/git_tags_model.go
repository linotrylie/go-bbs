package model

import (
	"fmt"
)

type GitTags struct {
	changes map[string]interface{}
	Tagid   int    `gorm:"primaryKey;column:tagid" json:"tagid"`
	Name    string `gorm:"column:name" json:"name"`
	Link    int    `gorm:"column:link" json:"link"`
}

func (*GitTags) TableName() string {
	return "bbs_git_tags"
}

// Location .
func (obj *GitTags) Location() map[string]interface{} {
	return map[string]interface{}{"tagid": obj.Tagid}
}

// Redis Key .
func (obj *GitTags) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Tagid)
}

// GetChanges .
func (obj *GitTags) GetChanges() map[string]interface{} {
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
func (obj *GitTags) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
func (obj *GitTags) SetTagid(val int) *GitTags {
	obj.Tagid = val
	obj.Update("tagid", obj.Tagid)
	return obj
}
func (obj *GitTags) SetName(val string) *GitTags {
	obj.Name = val
	obj.Update("name", obj.Name)
	return obj
}
func (obj *GitTags) SetLink(val int) *GitTags {
	obj.Link += val
	obj.Update("link", obj.Link)
	return obj
}
