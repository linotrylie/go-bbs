package model

type Tag_cate struct {
	changes      map[string]interface{}
	Cateid       int    `gorm:"primaryKey;column:cateid" json:"cateid"`
	Fid          int    `gorm:"column:fid" json:"fid"`
	Name         string `gorm:"column:name" json:"name"`
	Rank         int    `gorm:"column:rank" json:"rank"`
	Enable       int    `gorm:"column:enable" json:"enable"`
	Defaulttagid int    `gorm:"column:defaulttagid" json:"defaulttagid"`
	Isforce      int    `gorm:"column:isforce" json:"isforce"`
}

func (*Tag_cate) TableName() string {
	return "bbs_tag_cate"
}

// Location .
func (obj *Tag_cate) Location() map[string]interface{} {
	return map[string]interface{}{"cateid": obj.Cateid}
}

// GetChanges .
func (obj *Tag_cate) GetChanges() map[string]interface{} {
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
func (obj *Tag_cate) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
