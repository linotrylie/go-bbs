package model

import (
	"fmt"
)

type TableDay struct {
	changes    map[string]interface{}
	Year       int    `gorm:"primaryKey;column:year" json:"year"`    // 年
	Month      int    `gorm:"primaryKey;column:month" json:"month"`  // 月
	Day        int    `gorm:"primaryKey;column:day" json:"day"`      // 日
	CreateDate int64  `gorm:"column:create_date" json:"create_date"` // 时间戳
	Table      string `gorm:"primaryKey;column:table" json:"table"`  // 表名
	Maxid      int    `gorm:"column:maxid" json:"maxid"`             // 最大ID
	Count      int    `gorm:"column:count" json:"count"`             // 总数
}

func (*TableDay) TableName() string {
	return "bbs_table_day"
}

// Location .
func (obj *TableDay) Location() map[string]interface{} {
	return map[string]interface{}{"month": obj.Month, "day": obj.Day, "table": obj.Table, "year": obj.Year}
}

// Redis Key .
func (obj *TableDay) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Year) + "_" + fmt.Sprintf("%v", obj.Month) + "_" + fmt.Sprintf("%v", obj.Day) + "_" + fmt.Sprintf("%v", obj.Table)
}

// GetChanges .
func (obj *TableDay) GetChanges() map[string]interface{} {
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
func (obj *TableDay) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
func (obj *TableDay) SetYear(val int) *TableDay {
	obj.Year += val
	obj.Update("year", obj.Year)
	return obj
}
func (obj *TableDay) SetMonth(val int) *TableDay {
	obj.Month += val
	obj.Update("month", obj.Month)
	return obj
}
func (obj *TableDay) SetDay(val int) *TableDay {
	obj.Day += val
	obj.Update("day", obj.Day)
	return obj
}
func (obj *TableDay) SetCreateDate(val int64) *TableDay {
	obj.CreateDate += val
	obj.Update("create_date", obj.CreateDate)
	return obj
}
func (obj *TableDay) SetTable(val string) *TableDay {
	obj.Table = val
	obj.Update("table", obj.Table)
	return obj
}
func (obj *TableDay) SetMaxid(val int) *TableDay {
	obj.Maxid = val
	obj.Update("maxid", obj.Maxid)
	return obj
}
func (obj *TableDay) SetCount(val int) *TableDay {
	obj.Count += val
	obj.Update("count", obj.Count)
	return obj
}
