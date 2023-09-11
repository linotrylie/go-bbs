package model

import (
	"fmt"
)

type IqismartActivity struct {
	changes       map[string]interface{}
	Id            int    `gorm:"primaryKey;column:id" json:"id"`                // ID
	Name          string `gorm:"column:name" json:"name"`                       // 活动名称
	CreateDate    int    `gorm:"column:create_date" json:"create_date"`         // 创建时间
	StartDate     int    `gorm:"column:start_date" json:"start_date"`           // 开始时间
	EndDate       int    `gorm:"column:end_date" json:"end_date"`               // 结束时间
	ShowUserCount int    `gorm:"column:show_user_count" json:"show_user_count"` // 是否显示参数人数
	MaxUserCount  int    `gorm:"column:max_user_count" json:"max_user_count"`   // 参数人数限制
	Description   string `gorm:"column:description" json:"description"`         // 活动说明
	UseType       int    `gorm:"column:use_type" json:"use_type"`               // 消耗积分类型
	UseNum        int    `gorm:"column:use_num" json:"use_num"`                 // 消耗积分数量
	OnlyVip       int    `gorm:"column:only_vip" json:"only_vip"`               // 仅vip用户参与
	TimesLimit    int    `gorm:"column:times_limit" json:"times_limit"`         // 总参与次数限制
	TimesLimitDay int    `gorm:"column:times_limit_day" json:"times_limit_day"` // 每日参与次数限制
	Pid6          int    `gorm:"column:pid6" json:"pid6"`                       // 奖品6的商品ID
	Pid1          int    `gorm:"column:pid1" json:"pid1"`                       // 奖品1的商品ID
	Pid2          int    `gorm:"column:pid2" json:"pid2"`                       // 奖品2的商品ID
	Pid3          int    `gorm:"column:pid3" json:"pid3"`                       // 奖品3的商品ID
	Pid4          int    `gorm:"column:pid4" json:"pid4"`                       // 奖品4的商品ID
	Pid5          int    `gorm:"column:pid5" json:"pid5"`                       // 奖品5的商品ID
	Percent6      int    `gorm:"column:percent6" json:"percent6"`               // 6
	Percent1      int    `gorm:"column:percent1" json:"percent1"`               // 中奖率1
	Percent2      int    `gorm:"column:percent2" json:"percent2"`               // 中奖率2
	Percent3      int    `gorm:"column:percent3" json:"percent3"`               // 中奖率3
	Percent4      int    `gorm:"column:percent4" json:"percent4"`               // 中奖率4
	Percent5      int    `gorm:"column:percent5" json:"percent5"`               // 中奖率5
}

func (*IqismartActivity) TableName() string {
	return "bbs_iqismart_activity"
}

// Location .
func (obj *IqismartActivity) Location() map[string]interface{} {
	return map[string]interface{}{"id": obj.Id}
}

// Redis Key .
func (obj *IqismartActivity) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Id)
}

// GetChanges .
func (obj *IqismartActivity) GetChanges() map[string]interface{} {
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
func (obj *IqismartActivity) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
