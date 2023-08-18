package model

type ThreadTop struct {
	changes map[string]interface{}
	Fid     int `gorm:"column:fid" json:"fid"`
	Tid     int `gorm:"primaryKey;column:tid" json:"tid"`
	Top     int `gorm:"column:top" json:"top"`
}

func (*ThreadTop) TableName() string {
	return "bbs_thread_top"
}

// Location .
func (obj *ThreadTop) Location() map[string]interface{} {
	return map[string]interface{}{"tid": obj.Tid}
}

// GetChanges .
func (obj *ThreadTop) GetChanges() map[string]interface{} {
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
func (obj *ThreadTop) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
