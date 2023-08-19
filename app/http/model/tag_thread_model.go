package model

type TagThread struct {
	changes map[string]interface{}
	Tagid   int `gorm:"primaryKey;column:tagid" json:"tagid"`
	Tid     int `gorm:"primaryKey;column:tid" json:"tid"`
}

func (*TagThread) TableName() string {
	return "bbs_tag_thread"
}

// Location .
func (obj *TagThread) Location() map[string]interface{} {
	return map[string]interface{}{"tid": obj.Tid, "tagid": obj.Tagid}
}

// GetChanges .
func (obj *TagThread) GetChanges() map[string]interface{} {
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
func (obj *TagThread) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
