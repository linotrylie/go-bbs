package model

type Thread_top struct {
	changes map[string]interface{}
	Fid     int `gorm:"column:fid" json:"fid"`
	Tid     int `gorm:"primaryKey;column:tid" json:"tid"`
	Top     int `gorm:"column:top" json:"top"`
}

func (*Thread_top) TableName() string {
	return "bbs_thread_top"
}

// Location .
func (obj *Thread_top) Location() map[string]interface{} {
	return map[string]interface{}{"tid": obj.Tid}
}

// GetChanges .
func (obj *Thread_top) GetChanges() map[string]interface{} {
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
func (obj *Thread_top) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
