package model

type Table_day struct {
	changes    map[string]interface{}
	Year       int    `gorm:"primaryKey;column:year" json:"year"`   // 年
	Month      int    `gorm:"primaryKey;column:month" json:"month"` // 月
	Day        int    `gorm:"primaryKey;column:day" json:"day"`     // 日
	CreateDate int    `gorm:"column:createdate" json:"createdate"`  // 时间戳
	Table      string `gorm:"primaryKey;column:table" json:"table"` // 表名
	Maxid      int    `gorm:"column:maxid" json:"maxid"`            // 最大ID
	Count      int    `gorm:"column:count" json:"count"`            // 总数
}

func (*Table_day) TableName() string {
	return "bbs_table_day"
}

// Location .
func (obj *Table_day) Location() map[string]interface{} {
	return map[string]interface{}{"table": obj.Table, "year": obj.Year, "month": obj.Month, "day": obj.Day}
}

// GetChanges .
func (obj *Table_day) GetChanges() map[string]interface{} {
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
func (obj *Table_day) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
