package model

import (
	"fmt"
)

type Attach struct {
	changes     map[string]interface{}
	Aid         int    `gorm:"primaryKey;column:aid" json:"aid"`
	Tid         int    `gorm:"column:tid" json:"tid"`
	Pid         int    `gorm:"column:pid" json:"pid"`
	Uid         int    `gorm:"column:uid" json:"uid"`
	Filesize    int    `gorm:"column:filesize" json:"filesize"`
	Width       int    `gorm:"column:width" json:"width"`
	Height      int    `gorm:"column:height" json:"height"`
	Filename    string `gorm:"column:filename" json:"filename"`
	Orgfilename string `gorm:"column:orgfilename" json:"orgfilename"`
	Filetype    string `gorm:"column:filetype" json:"filetype"`
	CreateDate  int    `gorm:"column:create_date" json:"create_date"`
	Comment     string `gorm:"column:comment" json:"comment"`
	Downloads   int    `gorm:"column:downloads" json:"downloads"`
	Credits     int    `gorm:"column:credits" json:"credits"`
	Golds       int    `gorm:"column:golds" json:"golds"`
	Rmbs        int    `gorm:"column:rmbs" json:"rmbs"`
	Isimage     int    `gorm:"column:isimage" json:"isimage"`
}

func (*Attach) TableName() string {
	return "bbs_attach"
}

// Location .
func (obj *Attach) Location() map[string]interface{} {
	return map[string]interface{}{"aid": obj.Aid}
}

// Redis Key .
func (obj *Attach) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Aid)
}

// GetChanges .
func (obj *Attach) GetChanges() map[string]interface{} {
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
func (obj *Attach) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
