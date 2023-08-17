package model

type Sg_sign struct {
	changes      map[string]interface{}
	Id           int    `gorm:"column:id" json:"id"`                     // ID
	Uid          int    `gorm:"primaryKey;column:uid" json:"uid"`        // 用户ID
	Stime        int    `gorm:"column:stime" json:"stime"`               // 最后签到时间
	Credits      int    `gorm:"column:credits" json:"credits"`           // 签到积分
	Todaycredits int    `gorm:"column:todaycredits" json:"todaycredits"` // 今日积分
	Counts       int    `gorm:"column:counts" json:"counts"`             // 签到天数
	Keepdays     int    `gorm:"column:keepdays" json:"keepdays"`         // 连续签到
	User         string `gorm:"column:user" json:"user"`                 // 签到用户
}

func (*Sg_sign) TableName() string {
	return "sg_sign"
}

// Location .
func (obj *Sg_sign) Location() map[string]interface{} {
	return map[string]interface{}{"Uid": obj.Uid}
}

// GetChanges .
func (obj *Sg_sign) GetChanges() map[string]interface{} {
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
func (obj *Sg_sign) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
