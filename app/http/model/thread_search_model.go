package model

type ThreadSearch struct {
	changes map[string]interface{}
	Fid     int    `gorm:"column:fid" json:"fid"`
	Tid     int    `gorm:"primaryKey;column:tid" json:"tid"`
	Message string `gorm:"column:message" json:"message"`
}

func (*ThreadSearch) TableName() string {
	return "bbs_thread_search"
}

// Location .
func (obj *ThreadSearch) Location() map[string]interface{} {
	return map[string]interface{}{"tid": obj.Tid}
}

// GetChanges .
func (obj *ThreadSearch) GetChanges() map[string]interface{} {
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
func (obj *ThreadSearch) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
