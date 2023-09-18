package model

import (
	"fmt"
)

type Ipaccess struct {
	changes     map[string]interface{}
	Ip          uint32 `gorm:"primaryKey;column:ip" json:"ip"`
	Mails       int    `gorm:"column:mails" json:"mails"`
	Users       int    `gorm:"column:users" json:"users"`
	Logins      int    `gorm:"column:logins" json:"logins"`
	Threads     int    `gorm:"column:threads" json:"threads"`
	Posts       int    `gorm:"column:posts" json:"posts"`
	Attachs     int    `gorm:"column:attachs" json:"attachs"`
	Attachsizes int    `gorm:"column:attachsizes" json:"attachsizes"`
	LastDate    int64  `gorm:"column:last_date" json:"last_date"`
	Actions     int    `gorm:"column:actions" json:"actions"`
	Action1     int    `gorm:"column:action1" json:"action1"`
	Action2     int    `gorm:"column:action2" json:"action2"`
	Action3     int    `gorm:"column:action3" json:"action3"`
	Action4     int    `gorm:"column:action4" json:"action4"`
}

func (*Ipaccess) TableName() string {
	return "bbs_ipaccess"
}

// Location .
func (obj *Ipaccess) Location() map[string]interface{} {
	return map[string]interface{}{"ip": obj.Ip}
}

// Redis Key .
func (obj *Ipaccess) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Ip)
}

// GetChanges .
func (obj *Ipaccess) GetChanges() map[string]interface{} {
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
func (obj *Ipaccess) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
func (obj *Ipaccess) SetIp(val uint32) *Ipaccess {
	obj.Ip += val
	obj.Update("ip", obj.Ip)
	return obj
}
func (obj *Ipaccess) SetMails(val int) *Ipaccess {
	obj.Mails += val
	obj.Update("mails", obj.Mails)
	return obj
}
func (obj *Ipaccess) SetUsers(val int) *Ipaccess {
	obj.Users += val
	obj.Update("users", obj.Users)
	return obj
}
func (obj *Ipaccess) SetLogins(val int) *Ipaccess {
	obj.Logins += val
	obj.Update("logins", obj.Logins)
	return obj
}
func (obj *Ipaccess) SetThreads(val int) *Ipaccess {
	obj.Threads += val
	obj.Update("threads", obj.Threads)
	return obj
}
func (obj *Ipaccess) SetPosts(val int) *Ipaccess {
	obj.Posts += val
	obj.Update("posts", obj.Posts)
	return obj
}
func (obj *Ipaccess) SetAttachs(val int) *Ipaccess {
	obj.Attachs += val
	obj.Update("attachs", obj.Attachs)
	return obj
}
func (obj *Ipaccess) SetAttachsizes(val int) *Ipaccess {
	obj.Attachsizes += val
	obj.Update("attachsizes", obj.Attachsizes)
	return obj
}
func (obj *Ipaccess) SetLastDate(val int64) *Ipaccess {
	obj.LastDate += val
	obj.Update("last_date", obj.LastDate)
	return obj
}
func (obj *Ipaccess) SetActions(val int) *Ipaccess {
	obj.Actions += val
	obj.Update("actions", obj.Actions)
	return obj
}
func (obj *Ipaccess) SetAction1(val int) *Ipaccess {
	obj.Action1 += val
	obj.Update("action1", obj.Action1)
	return obj
}
func (obj *Ipaccess) SetAction2(val int) *Ipaccess {
	obj.Action2 += val
	obj.Update("action2", obj.Action2)
	return obj
}
func (obj *Ipaccess) SetAction3(val int) *Ipaccess {
	obj.Action3 += val
	obj.Update("action3", obj.Action3)
	return obj
}
func (obj *Ipaccess) SetAction4(val int) *Ipaccess {
	obj.Action4 += val
	obj.Update("action4", obj.Action4)
	return obj
}
