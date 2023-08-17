package model

type Tag struct {
	changes map[string]interface{}
	Tagid   int    `gorm:"primaryKey;column:tagid" json:"tagid"`
	Cateid  int    `gorm:"column:cateid" json:"cateid"`
	Name    string `gorm:"column:name" json:"name"`
	Rank    int    `gorm:"column:rank" json:"rank"`
	Enable  int    `gorm:"column:enable" json:"enable"`
	Style   string `gorm:"column:style" json:"style"`
}

func (*Tag) TableName() string {
	return "tag"
}

// Location .
func (obj *Tag) Location() map[string]interface{} {
	return map[string]interface{}{"Tagid": obj.Tagid}
}

// GetChanges .
func (obj *Tag) GetChanges() map[string]interface{} {
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
func (obj *Tag) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
