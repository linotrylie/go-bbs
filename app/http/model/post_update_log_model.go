package model

type Post_update_log struct {
	changes    map[string]interface{}
	Logid      int    `gorm:"primaryKey;column:logid" json:"logid"`
	Pid        int    `gorm:"column:pid" json:"pid"`
	Reason     string `gorm:"column:reason" json:"reason"`
	Message    string `gorm:"column:message" json:"message"`
	CreateDate int    `gorm:"column:createdate" json:"createdate"`
	Uid        int    `gorm:"column:uid" json:"uid"`
}

func (*Post_update_log) TableName() string {
	return "post_update_log"
}

// Location .
func (obj *Post_update_log) Location() map[string]interface{} {
	return map[string]interface{}{"Logid": obj.Logid}
}

// GetChanges .
func (obj *Post_update_log) GetChanges() map[string]interface{} {
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
func (obj *Post_update_log) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
