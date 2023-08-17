package model

type Git_tags struct {
	changes map[string]interface{}
	Tagid   int    `gorm:"primaryKey;column:tagid" json:"tagid"`
	Name    string `gorm:"column:name" json:"name"`
	Link    int    `gorm:"column:link" json:"link"`
}

func (*Git_tags) TableName() string {
	return "git_tags"
}

// Location .
func (obj *Git_tags) Location() map[string]interface{} {
	return map[string]interface{}{"Tagid": obj.Tagid}
}

// GetChanges .
func (obj *Git_tags) GetChanges() map[string]interface{} {
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
func (obj *Git_tags) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
