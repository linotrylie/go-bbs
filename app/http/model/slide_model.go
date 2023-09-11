package model

import (
	"fmt"
)

type Slide struct {
	changes   map[string]interface{}
	Slideid   int    `gorm:"primaryKey;column:slideid" json:"slideid"`
	Rank      int    `gorm:"column:rank" json:"rank"`
	Name      string `gorm:"column:name" json:"name"`
	Url       string `gorm:"column:url" json:"url"`
	Slidepic  string `gorm:"column:slidepic" json:"slidepic"`
	Picheight string `gorm:"column:picheight" json:"picheight"`
}

func (*Slide) TableName() string {
	return "bbs_slide"
}

// Location .
func (obj *Slide) Location() map[string]interface{} {
	return map[string]interface{}{"slideid": obj.Slideid}
}

// Redis Key .
func (obj *Slide) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Slideid)
}

// GetChanges .
func (obj *Slide) GetChanges() map[string]interface{} {
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
func (obj *Slide) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
