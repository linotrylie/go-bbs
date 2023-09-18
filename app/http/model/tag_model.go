package model

import (
	"fmt"
)

type Tag struct {
	changes map[string]interface{}
	Tagid   int    `gorm:"primaryKey;column:tagid" json:"tagid"`
	Cateid  int    `gorm:"column:cateid" json:"cateid"`
	Name    string `gorm:"column:name" json:"name"`
	Rank    int    `gorm:"column:rank" json:"rank"`
	Enable  int    `gorm:"column:enable" json:"enable"`
	Style   string `gorm:"column:style" json:"style"`
}

func (*Tag) TableName() string {
	return "bbs_tag"
}

// Location .
func (obj *Tag) Location() map[string]interface{} {
	return map[string]interface{}{"tagid": obj.Tagid}
}

// Redis Key .
func (obj *Tag) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Tagid)
}

// GetChanges .
func (obj *Tag) GetChanges() map[string]interface{} {
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
func (obj *Tag) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
func (obj *Tag) SetTagid(val int) *Tag {
	obj.Tagid = val
	obj.Update("tagid", obj.Tagid)
	return obj
}
func (obj *Tag) SetCateid(val int) *Tag {
	obj.Cateid = val
	obj.Update("cateid", obj.Cateid)
	return obj
}
func (obj *Tag) SetName(val string) *Tag {
	obj.Name = val
	obj.Update("name", obj.Name)
	return obj
}
func (obj *Tag) SetRank(val int) *Tag {
	obj.Rank += val
	obj.Update("rank", obj.Rank)
	return obj
}
func (obj *Tag) SetEnable(val int) *Tag {
	obj.Enable += val
	obj.Update("enable", obj.Enable)
	return obj
}
func (obj *Tag) SetStyle(val string) *Tag {
	obj.Style = val
	obj.Update("style", obj.Style)
	return obj
}
