package model

type Session struct {
	changes   map[string]interface{}
	Sid       string `gorm:"primaryKey;column:sid" json:"sid"`
	Uid       int    `gorm:"column:uid" json:"uid"`
	Fid       int    `gorm:"column:fid" json:"fid"`
	Url       string `gorm:"column:url" json:"url"`
	Ip        int    `gorm:"column:ip" json:"ip"`
	Useragent string `gorm:"column:useragent" json:"useragent"`
	Data      string `gorm:"column:data" json:"data"`
	Bigdata   int    `gorm:"column:bigdata" json:"bigdata"`
	LastDate  int    `gorm:"column:lastdate" json:"lastdate"`
}

func (*Session) TableName() string {
	return "session"
}

// Location .
func (obj *Session) Location() map[string]interface{} {
	return map[string]interface{}{"Sid": obj.Sid}
}

// GetChanges .
func (obj *Session) GetChanges() map[string]interface{} {
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
func (obj *Session) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
