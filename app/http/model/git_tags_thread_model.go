package model

type Git_tags_thread struct {
	changes map[string]interface{}
	Tagid   int `gorm:"primaryKey;column:tagid" json:"tagid"`
	Tid     int `gorm:"primaryKey;column:tid" json:"tid"`
}

func (*Git_tags_thread) TableName() string {
	return "git_tags_thread"
}

// Location .
func (obj *Git_tags_thread) Location() map[string]interface{} {
	return map[string]interface{}{"Tid": obj.Tid, "Tagid": obj.Tagid}
}

// GetChanges .
func (obj *Git_tags_thread) GetChanges() map[string]interface{} {
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
func (obj *Git_tags_thread) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
