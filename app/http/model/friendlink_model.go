package model

type Friendlink struct {
	changes    map[string]interface{}
	Linkid     int    `gorm:"primaryKey;column:linkid" json:"linkid"`
	Type       int    `gorm:"column:type" json:"type"`
	Rank       int    `gorm:"column:rank" json:"rank"`
	CreateDate int    `gorm:"column:createdate" json:"createdate"`
	Name       string `gorm:"column:name" json:"name"`
	Url        string `gorm:"column:url" json:"url"`
}

func (*Friendlink) TableName() string {
	return "bbs_friendlink"
}

// Location .
func (obj *Friendlink) Location() map[string]interface{} {
	return map[string]interface{}{"linkid": obj.Linkid}
}

// GetChanges .
func (obj *Friendlink) GetChanges() map[string]interface{} {
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
func (obj *Friendlink) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
