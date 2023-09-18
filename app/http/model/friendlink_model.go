package model

import (
	"fmt"
)

type Friendlink struct {
	changes    map[string]interface{}
	Linkid     int    `gorm:"primaryKey;column:linkid" json:"linkid"`
	Type       int    `gorm:"column:type" json:"type"`
	Rank       int    `gorm:"column:rank" json:"rank"`
	CreateDate int64  `gorm:"column:create_date" json:"create_date"`
	Name       string `gorm:"column:name" json:"name"`
	Url        string `gorm:"column:url" json:"url"`
}

func (*Friendlink) TableName() string {
	return "bbs_friendlink"
}

// Location .
func (obj *Friendlink) Location() map[string]interface{} {
	return map[string]interface{}{"linkid": obj.Linkid}
}

// Redis Key .
func (obj *Friendlink) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Linkid)
}

// GetChanges .
func (obj *Friendlink) GetChanges() map[string]interface{} {
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
func (obj *Friendlink) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
func (obj *Friendlink) SetLinkid(val int) *Friendlink {
	obj.Linkid = val
	obj.Update("linkid", obj.Linkid)
	return obj
}
func (obj *Friendlink) SetType(val int) *Friendlink {
	obj.Type += val
	obj.Update("type", obj.Type)
	return obj
}
func (obj *Friendlink) SetRank(val int) *Friendlink {
	obj.Rank += val
	obj.Update("rank", obj.Rank)
	return obj
}
func (obj *Friendlink) SetCreateDate(val int64) *Friendlink {
	obj.CreateDate += val
	obj.Update("create_date", obj.CreateDate)
	return obj
}
func (obj *Friendlink) SetName(val string) *Friendlink {
	obj.Name = val
	obj.Update("name", obj.Name)
	return obj
}
func (obj *Friendlink) SetUrl(val string) *Friendlink {
	obj.Url = val
	obj.Update("url", obj.Url)
	return obj
}
