package model

import (
	"fmt"
)

type GgFavoriteThread struct {
	changes map[string]interface{}
	Favid   int `gorm:"primaryKey;column:favid" json:"favid"`
	Tid     int `gorm:"column:tid" json:"tid"`
	Uid     int `gorm:"column:uid" json:"uid"`
}

func (*GgFavoriteThread) TableName() string {
	return "bbs_gg_favorite_thread"
}

// Location .
func (obj *GgFavoriteThread) Location() map[string]interface{} {
	return map[string]interface{}{"favid": obj.Favid}
}

// Redis Key .
func (obj *GgFavoriteThread) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Favid)
}

// GetChanges .
func (obj *GgFavoriteThread) GetChanges() map[string]interface{} {
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
func (obj *GgFavoriteThread) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
func (obj *GgFavoriteThread) SetFavid(val int) *GgFavoriteThread {
	obj.Favid = val
	obj.Update("favid", obj.Favid)
	return obj
}
func (obj *GgFavoriteThread) SetTid(val int) *GgFavoriteThread {
	obj.Tid = val
	obj.Update("tid", obj.Tid)
	return obj
}
func (obj *GgFavoriteThread) SetUid(val int) *GgFavoriteThread {
	obj.Uid = val
	obj.Update("uid", obj.Uid)
	return obj
}
