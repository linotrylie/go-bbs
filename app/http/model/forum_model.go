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
	CreateDate    int64  `gorm:"column:create_date" json:"create_date"`
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
func (obj *Forum) SetFid(val int) *Forum {
	obj.Fid = val
	obj.Update("fid", obj.Fid)
	return obj
}
func (obj *Forum) SetName(val string) *Forum {
	obj.Name = val
	obj.Update("name", obj.Name)
	return obj
}
func (obj *Forum) SetRank(val int) *Forum {
	obj.Rank += val
	obj.Update("rank", obj.Rank)
	return obj
}
func (obj *Forum) SetThreads(val int) *Forum {
	obj.Threads += val
	obj.Update("threads", obj.Threads)
	return obj
}
func (obj *Forum) SetTodayposts(val int) *Forum {
	obj.Todayposts += val
	obj.Update("todayposts", obj.Todayposts)
	return obj
}
func (obj *Forum) SetTodaythreads(val int) *Forum {
	obj.Todaythreads += val
	obj.Update("todaythreads", obj.Todaythreads)
	return obj
}
func (obj *Forum) SetBrief(val string) *Forum {
	obj.Brief = val
	obj.Update("brief", obj.Brief)
	return obj
}
func (obj *Forum) SetAnnouncement(val string) *Forum {
	obj.Announcement = val
	obj.Update("announcement", obj.Announcement)
	return obj
}
func (obj *Forum) SetAccesson(val int) *Forum {
	obj.Accesson += val
	obj.Update("accesson", obj.Accesson)
	return obj
}
func (obj *Forum) SetOrderby(val int) *Forum {
	obj.Orderby += val
	obj.Update("orderby", obj.Orderby)
	return obj
}
func (obj *Forum) SetCreateDate(val int64) *Forum {
	obj.CreateDate += val
	obj.Update("create_date", obj.CreateDate)
	return obj
}
func (obj *Forum) SetIcon(val int) *Forum {
	obj.Icon += val
	obj.Update("icon", obj.Icon)
	return obj
}
func (obj *Forum) SetModuids(val string) *Forum {
	obj.Moduids = val
	obj.Update("moduids", obj.Moduids)
	return obj
}
func (obj *Forum) SetSeoTitle(val string) *Forum {
	obj.SeoTitle = val
	obj.Update("seo_title", obj.SeoTitle)
	return obj
}
func (obj *Forum) SetSeoKeywords(val string) *Forum {
	obj.SeoKeywords = val
	obj.Update("seo_keywords", obj.SeoKeywords)
	return obj
}
func (obj *Forum) SetDigests(val int) *Forum {
	obj.Digests += val
	obj.Update("digests", obj.Digests)
	return obj
}
func (obj *Forum) SetCreateCredits(val int) *Forum {
	obj.CreateCredits += val
	obj.Update("create_credits", obj.CreateCredits)
	return obj
}
func (obj *Forum) SetCreateGolds(val int) *Forum {
	obj.CreateGolds += val
	obj.Update("create_golds", obj.CreateGolds)
	return obj
}
func (obj *Forum) SetPostCredits(val int) *Forum {
	obj.PostCredits += val
	obj.Update("post_credits", obj.PostCredits)
	return obj
}
func (obj *Forum) SetPostGolds(val int) *Forum {
	obj.PostGolds += val
	obj.Update("post_golds", obj.PostGolds)
	return obj
}
func (obj *Forum) SetAllowOffer(val int) *Forum {
	obj.AllowOffer += val
	obj.Update("allowoffer", obj.AllowOffer)
	return obj
}
