package model

type Mythread struct {
	changes map[string]interface{}
	Uid     int `gorm:"primaryKey;column:uid" json:"uid"`
	Tid     int `gorm:"primaryKey;column:tid" json:"tid"`
}

func (*Mythread) TableName() string {
	return "mythread"
}

// Location .
func (obj *Mythread) Location() map[string]interface{} {
	return map[string]interface{}{"Uid": obj.Uid, "Tid": obj.Tid}
}

// GetChanges .
func (obj *Mythread) GetChanges() map[string]interface{} {
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
func (obj *Mythread) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
