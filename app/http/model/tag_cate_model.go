package model

import (
	"fmt"
)

type TagCate struct {
	changes      map[string]interface{}
	Cateid       int    `gorm:"primaryKey;column:cateid" json:"cateid"`
	Fid          int    `gorm:"column:fid" json:"fid"`
	Name         string `gorm:"column:name" json:"name"`
	Rank         int    `gorm:"column:rank" json:"rank"`
	Enable       int    `gorm:"column:enable" json:"enable"`
	Defaulttagid int    `gorm:"column:defaulttagid" json:"defaulttagid"`
	Isforce      int    `gorm:"column:isforce" json:"isforce"`
}

func (*TagCate) TableName() string {
	return "bbs_tag_cate"
}

// Location .
func (obj *TagCate) Location() map[string]interface{} {
	return map[string]interface{}{"cateid": obj.Cateid}
}

// Redis Key .
func (obj *TagCate) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Cateid)
}

// GetChanges .
func (obj *TagCate) GetChanges() map[string]interface{} {
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
func (obj *TagCate) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
func (obj *TagCate) SetCateid(val int) *TagCate {
	obj.Cateid = val
	obj.Update("cateid", obj.Cateid)
	return obj
}
func (obj *TagCate) SetFid(val int) *TagCate {
	obj.Fid = val
	obj.Update("fid", obj.Fid)
	return obj
}
func (obj *TagCate) SetName(val string) *TagCate {
	obj.Name = val
	obj.Update("name", obj.Name)
	return obj
}
func (obj *TagCate) SetRank(val int) *TagCate {
	obj.Rank += val
	obj.Update("rank", obj.Rank)
	return obj
}
func (obj *TagCate) SetEnable(val int) *TagCate {
	obj.Enable += val
	obj.Update("enable", obj.Enable)
	return obj
}
func (obj *TagCate) SetDefaulttagid(val int) *TagCate {
	obj.Defaulttagid = val
	obj.Update("defaulttagid", obj.Defaulttagid)
	return obj
}
func (obj *TagCate) SetIsforce(val int) *TagCate {
	obj.Isforce += val
	obj.Update("isforce", obj.Isforce)
	return obj
}
