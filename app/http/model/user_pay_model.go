package model

type User_pay struct {
	changes    map[string]interface{}
	Cid        int    `gorm:"primaryKey;column:cid" json:"cid"`
	Uid        int    `gorm:"column:uid" json:"uid"`
	Status     int    `gorm:"column:status" json:"status"`
	Num        int    `gorm:"column:num" json:"num"`
	Type       int    `gorm:"column:type" json:"type"`
	CreditType int    `gorm:"column:credittype" json:"credittype"`
	Code       string `gorm:"column:code" json:"code"`
	Time       int    `gorm:"column:time" json:"time"`
}

func (*User_pay) TableName() string {
	return "bbs_user_pay"
}

// Location .
func (obj *User_pay) Location() map[string]interface{} {
	return map[string]interface{}{"cid": obj.Cid}
}

// GetChanges .
func (obj *User_pay) GetChanges() map[string]interface{} {
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
func (obj *User_pay) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
