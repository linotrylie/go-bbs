package model

type Thread_search struct {
	changes map[string]interface{}
	Fid     int    `gorm:"column:fid" json:"fid"`
	Tid     int    `gorm:"primaryKey;column:tid" json:"tid"`
	Message string `gorm:"column:message" json:"message"`
}

func (*Thread_search) TableName() string {
	return "bbs_thread_search"
}

// Location .
func (obj *Thread_search) Location() map[string]interface{} {
	return map[string]interface{}{"tid": obj.Tid}
}

// GetChanges .
func (obj *Thread_search) GetChanges() map[string]interface{} {
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
func (obj *Thread_search) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
