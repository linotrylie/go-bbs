package model

import (
	"fmt"
)

type Ipaccess struct {
	changes     map[string]interface{}
	Ip          int `gorm:"primaryKey;column:ip" json:"ip"`
	Mails       int `gorm:"column:mails" json:"mails"`
	Users       int `gorm:"column:users" json:"users"`
	Logins      int `gorm:"column:logins" json:"logins"`
	Threads     int `gorm:"column:threads" json:"threads"`
	Posts       int `gorm:"column:posts" json:"posts"`
	Attachs     int `gorm:"column:attachs" json:"attachs"`
	Attachsizes int `gorm:"column:attachsizes" json:"attachsizes"`
	LastDate    int `gorm:"column:last_date" json:"last_date"`
	Actions     int `gorm:"column:actions" json:"actions"`
	Action1     int `gorm:"column:action1" json:"action1"`
	Action2     int `gorm:"column:action2" json:"action2"`
	Action3     int `gorm:"column:action3" json:"action3"`
	Action4     int `gorm:"column:action4" json:"action4"`
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
