package model

type Session_data struct {
	changes  map[string]interface{}
	Sid      string `gorm:"primaryKey;column:sid" json:"sid"`
	LastDate int    `gorm:"column:lastdate" json:"lastdate"`
	Data     string `gorm:"column:data" json:"data"`
}

func (*Session_data) TableName() string {
	return "session_data"
}

// Location .
func (obj *Session_data) Location() map[string]interface{} {
	return map[string]interface{}{"Sid": obj.Sid}
}

// GetChanges .
func (obj *Session_data) GetChanges() map[string]interface{} {
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
func (obj *Session_data) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
