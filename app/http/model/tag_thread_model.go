package model

type Tag_thread struct {
	changes map[string]interface{}
	Tagid   int `gorm:"primaryKey;column:tagid" json:"tagid"`
	Tid     int `gorm:"primaryKey;column:tid" json:"tid"`
}

func (*Tag_thread) TableName() string {
	return "bbs_tag_thread"
}

// Location .
func (obj *Tag_thread) Location() map[string]interface{} {
	return map[string]interface{}{"tid": obj.Tid, "tagid": obj.Tagid}
}

// GetChanges .
func (obj *Tag_thread) GetChanges() map[string]interface{} {
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
func (obj *Tag_thread) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
