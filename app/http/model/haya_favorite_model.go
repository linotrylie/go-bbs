package model

import (
	"fmt"
	"time"
)

type HayaFavorite struct {
	changes    map[string]interface{}
	Tid        int    `gorm:"column:tid" json:"tid"`                 // 帖子ID
	Uid        int    `gorm:"column:uid" json:"uid"`                 // 用户ID
	CreateDate int64  `gorm:"column:create_date" json:"create_date"` // 添加时间
	CreateIp   uint32 `gorm:"column:create_ip" json:"create_ip"`     // 添加IP
}

func (*HayaFavorite) TableName() string {
	return "bbs_haya_favorite"
}

// Location .
func (obj *HayaFavorite) Location() map[string]interface{} {
	return map[string]interface{}{}
}

// Redis Key .
func (obj *HayaFavorite) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", time.Now().Unix())
}

// GetChanges .
func (obj *HayaFavorite) GetChanges() map[string]interface{} {
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
func (obj *HayaFavorite) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
func (obj *HayaFavorite) SetTid(val int) *HayaFavorite {
	obj.Tid = val
	obj.Update("tid", obj.Tid)
	return obj
}
func (obj *HayaFavorite) SetUid(val int) *HayaFavorite {
	obj.Uid = val
	obj.Update("uid", obj.Uid)
	return obj
}
func (obj *HayaFavorite) SetCreateDate(val int64) *HayaFavorite {
	obj.CreateDate += val
	obj.Update("create_date", obj.CreateDate)
	return obj
}
func (obj *HayaFavorite) SetCreateIp(val uint32) *HayaFavorite {
	obj.CreateIp += val
	obj.Update("create_ip", obj.CreateIp)
	return obj
}
