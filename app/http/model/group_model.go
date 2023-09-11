package model

import (
	"fmt"
)

type Group struct {
	changes         map[string]interface{}
	Gid             int    `gorm:"primaryKey;column:gid" json:"gid"`
	Name            string `gorm:"column:name" json:"name"`
	Creditsfrom     int    `gorm:"column:creditsfrom" json:"creditsfrom"`
	Creditsto       int    `gorm:"column:creditsto" json:"creditsto"`
	Allowread       int    `gorm:"column:allowread" json:"allowread"`
	Allowthread     int    `gorm:"column:allowthread" json:"allowthread"`
	Allowpost       int    `gorm:"column:allowpost" json:"allowpost"`
	Allowattach     int    `gorm:"column:allowattach" json:"allowattach"`
	Allowdown       int    `gorm:"column:allowdown" json:"allowdown"`
	Allowtop        int    `gorm:"column:allowtop" json:"allowtop"`
	Allowupdate     int    `gorm:"column:allowupdate" json:"allowupdate"`
	Allowdelete     int    `gorm:"column:allowdelete" json:"allowdelete"`
	Allowmove       int    `gorm:"column:allowmove" json:"allowmove"`
	Allowbanuser    int    `gorm:"column:allowbanuser" json:"allowbanuser"`
	Allowdeleteuser int    `gorm:"column:allowdeleteuser" json:"allowdeleteuser"`
	Allowviewip     int    `gorm:"column:allowviewip" json:"allowviewip"`
	Allowharddelete int    `gorm:"column:allowharddelete" json:"allowharddelete"`
	Readp           int    `gorm:"column:readp" json:"readp"`
	AllowPostRead   int    `gorm:"column:allowpostread" json:"allowpostread"`
	Allowsell       int    `gorm:"column:allowsell" json:"allowsell"`
	AllowOffer      int    `gorm:"column:allowoffer" json:"allowoffer"`
}

func (*Group) TableName() string {
	return "bbs_group"
}

// Location .
func (obj *Group) Location() map[string]interface{} {
	return map[string]interface{}{"gid": obj.Gid}
}

// Redis Key .
func (obj *Group) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Gid)
}

// GetChanges .
func (obj *Group) GetChanges() map[string]interface{} {
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
func (obj *Group) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
