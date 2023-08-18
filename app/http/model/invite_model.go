package model

type Invite struct {
	changes map[string]interface{}
	Uid     int `gorm:"primaryKey;column:uid" json:"uid"`
	Ip      int `gorm:"column:ip" json:"ip"`
	Regtime int `gorm:"column:regtime" json:"regtime"`
}

func (*Invite) TableName() string {
	return "bbs_invite"
}

// Location .
func (obj *Invite) Location() map[string]interface{} {
	return map[string]interface{}{"uid": obj.Uid}
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
