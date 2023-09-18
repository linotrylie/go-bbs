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
	Allowviewip     uint32 `gorm:"column:allowviewip" json:"allowviewip"`
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
func (obj *Group) SetGid(val int) *Group {
	obj.Gid = val
	obj.Update("gid", obj.Gid)
	return obj
}
func (obj *Group) SetName(val string) *Group {
	obj.Name = val
	obj.Update("name", obj.Name)
	return obj
}
func (obj *Group) SetCreditsfrom(val int) *Group {
	obj.Creditsfrom += val
	obj.Update("creditsfrom", obj.Creditsfrom)
	return obj
}
func (obj *Group) SetCreditsto(val int) *Group {
	obj.Creditsto += val
	obj.Update("creditsto", obj.Creditsto)
	return obj
}
func (obj *Group) SetAllowread(val int) *Group {
	obj.Allowread += val
	obj.Update("allowread", obj.Allowread)
	return obj
}
func (obj *Group) SetAllowthread(val int) *Group {
	obj.Allowthread += val
	obj.Update("allowthread", obj.Allowthread)
	return obj
}
func (obj *Group) SetAllowpost(val int) *Group {
	obj.Allowpost += val
	obj.Update("allowpost", obj.Allowpost)
	return obj
}
func (obj *Group) SetAllowattach(val int) *Group {
	obj.Allowattach += val
	obj.Update("allowattach", obj.Allowattach)
	return obj
}
func (obj *Group) SetAllowdown(val int) *Group {
	obj.Allowdown += val
	obj.Update("allowdown", obj.Allowdown)
	return obj
}
func (obj *Group) SetAllowtop(val int) *Group {
	obj.Allowtop += val
	obj.Update("allowtop", obj.Allowtop)
	return obj
}
func (obj *Group) SetAllowupdate(val int) *Group {
	obj.Allowupdate += val
	obj.Update("allowupdate", obj.Allowupdate)
	return obj
}
func (obj *Group) SetAllowdelete(val int) *Group {
	obj.Allowdelete += val
	obj.Update("allowdelete", obj.Allowdelete)
	return obj
}
func (obj *Group) SetAllowmove(val int) *Group {
	obj.Allowmove += val
	obj.Update("allowmove", obj.Allowmove)
	return obj
}
func (obj *Group) SetAllowbanuser(val int) *Group {
	obj.Allowbanuser += val
	obj.Update("allowbanuser", obj.Allowbanuser)
	return obj
}
func (obj *Group) SetAllowdeleteuser(val int) *Group {
	obj.Allowdeleteuser += val
	obj.Update("allowdeleteuser", obj.Allowdeleteuser)
	return obj
}
func (obj *Group) SetAllowviewip(val uint32) *Group {
	obj.Allowviewip += val
	obj.Update("allowviewip", obj.Allowviewip)
	return obj
}
func (obj *Group) SetAllowharddelete(val int) *Group {
	obj.Allowharddelete += val
	obj.Update("allowharddelete", obj.Allowharddelete)
	return obj
}
func (obj *Group) SetReadp(val int) *Group {
	obj.Readp += val
	obj.Update("readp", obj.Readp)
	return obj
}
func (obj *Group) SetAllowPostRead(val int) *Group {
	obj.AllowPostRead += val
	obj.Update("allowpostread", obj.AllowPostRead)
	return obj
}
func (obj *Group) SetAllowsell(val int) *Group {
	obj.Allowsell += val
	obj.Update("allowsell", obj.Allowsell)
	return obj
}
func (obj *Group) SetAllowOffer(val int) *Group {
	obj.AllowOffer += val
	obj.Update("allowoffer", obj.AllowOffer)
	return obj
}
