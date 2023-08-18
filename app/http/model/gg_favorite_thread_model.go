package model

type Gg_favorite_thread struct {
	changes map[string]interface{}
	Favid   int `gorm:"primaryKey;column:favid" json:"favid"`
	Tid     int `gorm:"column:tid" json:"tid"`
	Uid     int `gorm:"column:uid" json:"uid"`
}

func (*Gg_favorite_thread) TableName() string {
	return "bbs_gg_favorite_thread"
}

// Location .
func (obj *Gg_favorite_thread) Location() map[string]interface{} {
	return map[string]interface{}{"favid": obj.Favid}
}

// GetChanges .
func (obj *Gg_favorite_thread) GetChanges() map[string]interface{} {
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
func (obj *Gg_favorite_thread) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
