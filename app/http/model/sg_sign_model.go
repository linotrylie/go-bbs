package model

import (
	"fmt"
)

type SgSign struct {
	changes      map[string]interface{}
	Id           int    `gorm:"column:id" json:"id"`                     // ID
	Uid          int    `gorm:"primaryKey;column:uid" json:"uid"`        // 用户ID
	Stime        int64  `gorm:"column:stime" json:"stime"`               // 最后签到时间
	Credits      int    `gorm:"column:credits" json:"credits"`           // 签到积分
	Todaycredits int    `gorm:"column:todaycredits" json:"todaycredits"` // 今日积分
	Counts       int    `gorm:"column:counts" json:"counts"`             // 签到天数
	Keepdays     int    `gorm:"column:keepdays" json:"keepdays"`         // 连续签到
	User         string `gorm:"column:user" json:"user"`                 // 签到用户
}

func (*SgSign) TableName() string {
	return "bbs_sg_sign"
}

// Location .
func (obj *SgSign) Location() map[string]interface{} {
	return map[string]interface{}{"uid": obj.Uid}
}

// Redis Key .
func (obj *SgSign) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Uid)
}

// GetChanges .
func (obj *SgSign) GetChanges() map[string]interface{} {
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
func (obj *SgSign) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
func (obj *SgSign) SetId(val int) *SgSign {
	obj.Id += val
	obj.Update("id", obj.Id)
	return obj
}
func (obj *SgSign) SetUid(val int) *SgSign {
	obj.Uid = val
	obj.Update("uid", obj.Uid)
	return obj
}
func (obj *SgSign) SetStime(val int64) *SgSign {
	obj.Stime += val
	obj.Update("stime", obj.Stime)
	return obj
}
func (obj *SgSign) SetCredits(val int) *SgSign {
	obj.Credits += val
	obj.Update("credits", obj.Credits)
	return obj
}
func (obj *SgSign) SetTodaycredits(val int) *SgSign {
	obj.Todaycredits += val
	obj.Update("todaycredits", obj.Todaycredits)
	return obj
}
func (obj *SgSign) SetCounts(val int) *SgSign {
	obj.Counts += val
	obj.Update("counts", obj.Counts)
	return obj
}
func (obj *SgSign) SetKeepdays(val int) *SgSign {
	obj.Keepdays += val
	obj.Update("keepdays", obj.Keepdays)
	return obj
}
func (obj *SgSign) SetUser(val string) *SgSign {
	obj.User = val
	obj.Update("user", obj.User)
	return obj
}
