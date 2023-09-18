package model

import (
	"fmt"
)

type Invite struct {
	changes map[string]interface{}
	Uid     int    `gorm:"primaryKey;column:uid" json:"uid"`
	Ip      uint32 `gorm:"column:ip" json:"ip"`
	Regtime int64  `gorm:"column:regtime" json:"regtime"`
}

func (*Invite) TableName() string {
	return "bbs_invite"
}

// Location .
func (obj *Invite) Location() map[string]interface{} {
	return map[string]interface{}{"uid": obj.Uid}
}

// Redis Key .
func (obj *Invite) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Uid)
}

// GetChanges .
func (obj *Invite) GetChanges() map[string]interface{} {
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
func (obj *Invite) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
func (obj *Invite) SetUid(val int) *Invite {
	obj.Uid = val
	obj.Update("uid", obj.Uid)
	return obj
}
func (obj *Invite) SetIp(val uint32) *Invite {
	obj.Ip += val
	obj.Update("ip", obj.Ip)
	return obj
}
func (obj *Invite) SetRegtime(val int64) *Invite {
	obj.Regtime += val
	obj.Update("regtime", obj.Regtime)
	return obj
}
