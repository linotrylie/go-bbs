package model

import (
	"fmt"
)

type IqismartProduct struct {
	changes    map[string]interface{}
	Id         int    `gorm:"primaryKey;column:id" json:"id"`        // ID
	Name       string `gorm:"column:name" json:"name"`               // 商品名称
	Type       int    `gorm:"column:type" json:"type"`               // 商品类型：0金币 1人民币 2主题 3vip会员 4实物
	Code       string `gorm:"column:code" json:"code"`               // 数量、主题ID、会员月份
	CreateDate int    `gorm:"column:create_date" json:"create_date"` // 创建时间
}

func (*IqismartProduct) TableName() string {
	return "bbs_iqismart_product"
}

// Location .
func (obj *IqismartProduct) Location() map[string]interface{} {
	return map[string]interface{}{"id": obj.Id}
}

// Redis Key .
func (obj *IqismartProduct) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Id)
}

// GetChanges .
func (obj *IqismartProduct) GetChanges() map[string]interface{} {
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
func (obj *IqismartProduct) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
