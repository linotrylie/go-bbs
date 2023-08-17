package model

type Modlog struct {
	changes    map[string]interface{}
	Logid      int    `gorm:"primaryKey;column:logid" json:"logid"`
	Uid        int    `gorm:"column:uid" json:"uid"`
	Tid        int    `gorm:"column:tid" json:"tid"`
	Pid        int    `gorm:"column:pid" json:"pid"`
	Subject    string `gorm:"column:subject" json:"subject"`
	Comment    string `gorm:"column:comment" json:"comment"`
	Rmbs       int    `gorm:"column:rmbs" json:"rmbs"`
	CreateDate int    `gorm:"column:createdate" json:"createdate"`
	Action     string `gorm:"column:action" json:"action"`
}

func (*Modlog) TableName() string {
	return "modlog"
}

// Location .
func (obj *Modlog) Location() map[string]interface{} {
	return map[string]interface{}{"Logid": obj.Logid}
}

// GetChanges .
func (obj *Modlog) GetChanges() map[string]interface{} {
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
func (obj *Modlog) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
