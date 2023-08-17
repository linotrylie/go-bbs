package model

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
	CreateDate    int    `gorm:"column:createdate" json:"createdate"`
	Icon          int    `gorm:"column:icon" json:"icon"`
	Moduids       string `gorm:"column:moduids" json:"moduids"`
	SeoTitle      string `gorm:"column:seotitle" json:"seotitle"`
	SeoKeywords   string `gorm:"column:seokeywords" json:"seokeywords"`
	Digests       int    `gorm:"column:digests" json:"digests"`
	CreateCredits int    `gorm:"column:createcredits" json:"createcredits"`
	CreateGolds   int    `gorm:"column:creategolds" json:"creategolds"`
	PostCredits   int    `gorm:"column:postcredits" json:"postcredits"`
	PostGolds     int    `gorm:"column:postgolds" json:"postgolds"`
	AllowOffer    int    `gorm:"column:allowoffer" json:"allowoffer"`
}

func (*Forum) TableName() string {
	return "forum"
}

// Location .
func (obj *Forum) Location() map[string]interface{} {
	return map[string]interface{}{"Fid": obj.Fid}
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
