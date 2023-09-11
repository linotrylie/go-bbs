package model

import (
	"fmt"
)

type IqismartActivityOrder struct {
	changes    map[string]interface{}
	Id         int    `gorm:"primaryKey;column:id" json:"id"`        // ID
	Uid        int    `gorm:"column:uid" json:"uid"`                 // uid
	Aid        int    `gorm:"column:aid" json:"aid"`                 // 活动iD
	Pid        int    `gorm:"column:pid" json:"pid"`                 // 商品ID
	Code       string `gorm:"column:code" json:"code"`               // 数量、主题ID、会员月份、实物兑换码
	CreateDate int    `gorm:"column:create_date" json:"create_date"` // 创建时间
}

func (*IqismartActivityOrder) TableName() string {
	return "bbs_iqismart_activity_order"
}

// Location .
func (obj *IqismartActivityOrder) Location() map[string]interface{} {
	return map[string]interface{}{"id": obj.Id}
}

// Redis Key .
func (obj *IqismartActivityOrder) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Id)
}

// GetChanges .
func (obj *IqismartActivityOrder) GetChanges() map[string]interface{} {
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
func (obj *IqismartActivityOrder) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
