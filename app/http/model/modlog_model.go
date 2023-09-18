package model

import (
	"fmt"
)

type Modlog struct {
	changes    map[string]interface{}
	Logid      int    `gorm:"primaryKey;column:logid" json:"logid"`
	Uid        int    `gorm:"column:uid" json:"uid"`
	Tid        int    `gorm:"column:tid" json:"tid"`
	Pid        int    `gorm:"column:pid" json:"pid"`
	Subject    string `gorm:"column:subject" json:"subject"`
	Comment    string `gorm:"column:comment" json:"comment"`
	Rmbs       int    `gorm:"column:rmbs" json:"rmbs"`
	CreateDate int64  `gorm:"column:create_date" json:"create_date"`
	Action     string `gorm:"column:action" json:"action"`
}

func (*Modlog) TableName() string {
	return "bbs_modlog"
}

// Location .
func (obj *Modlog) Location() map[string]interface{} {
	return map[string]interface{}{"logid": obj.Logid}
}

// Redis Key .
func (obj *Modlog) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Logid)
}

// GetChanges .
func (obj *Modlog) GetChanges() map[string]interface{} {
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
func (obj *Modlog) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
func (obj *Modlog) SetLogid(val int) *Modlog {
	obj.Logid = val
	obj.Update("logid", obj.Logid)
	return obj
}
func (obj *Modlog) SetUid(val int) *Modlog {
	obj.Uid = val
	obj.Update("uid", obj.Uid)
	return obj
}
func (obj *Modlog) SetTid(val int) *Modlog {
	obj.Tid = val
	obj.Update("tid", obj.Tid)
	return obj
}
func (obj *Modlog) SetPid(val int) *Modlog {
	obj.Pid = val
	obj.Update("pid", obj.Pid)
	return obj
}
func (obj *Modlog) SetSubject(val string) *Modlog {
	obj.Subject = val
	obj.Update("subject", obj.Subject)
	return obj
}
func (obj *Modlog) SetComment(val string) *Modlog {
	obj.Comment = val
	obj.Update("comment", obj.Comment)
	return obj
}
func (obj *Modlog) SetRmbs(val int) *Modlog {
	obj.Rmbs += val
	obj.Update("rmbs", obj.Rmbs)
	return obj
}
func (obj *Modlog) SetCreateDate(val int64) *Modlog {
	obj.CreateDate += val
	obj.Update("create_date", obj.CreateDate)
	return obj
}
func (obj *Modlog) SetAction(val string) *Modlog {
	obj.Action = val
	obj.Update("action", obj.Action)
	return obj
}
