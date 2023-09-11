package model

import (
	"fmt"
)

type Forum struct {
	changes       map[string]interface{}
	Fid           int    `gorm:"primaryKey;column:fid" json:"fid"`
	Name          string `gorm:"column:name" json:"name"`
	Rank          int    `gorm:"column:rank" json:"rank"`
	Threads       int    `gorm:"column:threads" json:"threads"`
	Todayposts    int    `gorm:"column:todayposts" json:"todayposts"`
	Todaythreads  int    `gorm:"column:todaythreads" json:"todaythreads"`
	Brief         string `gorm:"column:brief" json:"brief"`
	Announcement  string `gorm:"column:announcement" json:"announcement"`
	Accesson      int    `gorm:"column:accesson" json:"accesson"`
	Orderby       int    `gorm:"column:orderby" json:"orderby"`
	CreateDate    int    `gorm:"column:create_date" json:"create_date"`
	Icon          int    `gorm:"column:icon" json:"icon"`
	Moduids       string `gorm:"column:moduids" json:"moduids"`
	SeoTitle      string `gorm:"column:seo_title" json:"seo_title"`
	SeoKeywords   string `gorm:"column:seo_keywords" json:"seo_keywords"`
	Digests       int    `gorm:"column:digests" json:"digests"`
	CreateCredits int    `gorm:"column:create_credits" json:"create_credits"`
	CreateGolds   int    `gorm:"column:create_golds" json:"create_golds"`
	PostCredits   int    `gorm:"column:post_credits" json:"post_credits"`
	PostGolds     int    `gorm:"column:post_golds" json:"post_golds"`
	AllowOffer    int    `gorm:"column:allowoffer" json:"allowoffer"`
}

func (*Forum) TableName() string {
	return "bbs_forum"
}

// Location .
func (obj *Forum) Location() map[string]interface{} {
	return map[string]interface{}{"fid": obj.Fid}
}

// Redis Key .
func (obj *Forum) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Fid)
}

// GetChanges .
func (obj *Forum) GetChanges() map[string]interface{} {
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
func (obj *Forum) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
