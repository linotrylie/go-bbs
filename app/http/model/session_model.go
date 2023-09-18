package model

import (
	"fmt"
)

type Session struct {
	changes   map[string]interface{}
	Sid       string `gorm:"primaryKey;column:sid" json:"sid"`
	Uid       int    `gorm:"column:uid" json:"uid"`
	Fid       int    `gorm:"column:fid" json:"fid"`
	Url       string `gorm:"column:url" json:"url"`
	Ip        uint32 `gorm:"column:ip" json:"ip"`
	Useragent string `gorm:"column:useragent" json:"useragent"`
	Data      string `gorm:"column:data" json:"data"`
	Bigdata   int    `gorm:"column:bigdata" json:"bigdata"`
	LastDate  int64  `gorm:"column:last_date" json:"last_date"`
}

func (*Session) TableName() string {
	return "bbs_session"
}

// Location .
func (obj *Session) Location() map[string]interface{} {
	return map[string]interface{}{"sid": obj.Sid}
}

// Redis Key .
func (obj *Session) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Sid)
}

// GetChanges .
func (obj *Session) GetChanges() map[string]interface{} {
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
func (obj *Session) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
func (obj *Session) SetSid(val string) *Session {
	obj.Sid = val
	obj.Update("sid", obj.Sid)
	return obj
}
func (obj *Session) SetUid(val int) *Session {
	obj.Uid = val
	obj.Update("uid", obj.Uid)
	return obj
}
func (obj *Session) SetFid(val int) *Session {
	obj.Fid = val
	obj.Update("fid", obj.Fid)
	return obj
}
func (obj *Session) SetUrl(val string) *Session {
	obj.Url = val
	obj.Update("url", obj.Url)
	return obj
}
func (obj *Session) SetIp(val uint32) *Session {
	obj.Ip += val
	obj.Update("ip", obj.Ip)
	return obj
}
func (obj *Session) SetUseragent(val string) *Session {
	obj.Useragent = val
	obj.Update("useragent", obj.Useragent)
	return obj
}
func (obj *Session) SetData(val string) *Session {
	obj.Data = val
	obj.Update("data", obj.Data)
	return obj
}
func (obj *Session) SetBigdata(val int) *Session {
	obj.Bigdata += val
	obj.Update("bigdata", obj.Bigdata)
	return obj
}
func (obj *Session) SetLastDate(val int64) *Session {
	obj.LastDate += val
	obj.Update("last_date", obj.LastDate)
	return obj
}
