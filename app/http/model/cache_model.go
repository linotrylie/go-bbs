package model

type Cache struct {
	changes map[string]interface{}
	K       string `gorm:"primaryKey;column:k" json:"k"`
	V       string `gorm:"column:v" json:"v"`
	Expiry  int    `gorm:"column:expiry" json:"expiry"`
}

func (*Cache) TableName() string {
	return "bbs_cache"
}

// Location .
func (obj *Cache) Location() map[string]interface{} {
	return map[string]interface{}{"k": obj.K}
}

// GetChanges .
func (obj *Cache) GetChanges() map[string]interface{} {
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
func (obj *Cache) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
