package model

import (
	"fmt"
)

type Post struct {
	changes          map[string]interface{}
	Tid              int      `gorm:"column:tid" json:"tid"`
	Pid              int      `gorm:"primaryKey;column:pid" json:"pid"`
	Uid              int      `gorm:"column:uid" json:"uid"`
	Isfirst          int      `gorm:"column:isfirst" json:"isfirst"`
	CreateDate       int64    `gorm:"column:create_date" json:"create_date"`
	Userip           uint32   `gorm:"column:userip" json:"userip"`
	Images           int      `gorm:"column:images" json:"images"`
	Files            int      `gorm:"column:files" json:"files"`
	Doctype          int      `gorm:"column:doctype" json:"doctype"`
	Quotepid         int      `gorm:"column:quotepid" json:"quotepid"`
	Message          string   `gorm:"column:message" json:"message"`
	MessageFmt       string   `gorm:"column:message_fmt" json:"message_fmt"`
	LocationPost     string   `gorm:"column:location_post" json:"location_post"`
	Likes            int      `gorm:"column:likes" json:"likes"` // 点赞数
	Deleted          int      `gorm:"column:deleted" json:"deleted"`
	Updates          int      `gorm:"column:updates" json:"updates"`
	LastUpdateDate   int64    `gorm:"column:last_update_date" json:"last_update_date"`
	LastUpdateUid    int      `gorm:"column:last_update_uid" json:"last_update_uid"`
	LastUpdateReason string   `gorm:"column:last_update_reason" json:"last_update_reason"`
	ReplyHide        int      `gorm:"column:reply_hide" json:"reply_hide"`
	Attach           []Attach `gorm:"foreignKey:pid;references:pid" json:"-"`
	CreateUser       User     `gorm:"foreignKey:uid;references:uid" json:"-"`
	LastUpdateUser   User     `gorm:"foreignKey:last_update_uid;references:last_update_uid" json:"-"`
}

func (*Post) TableName() string {
	return "bbs_post"
}

// Location .
func (obj *Post) Location() map[string]interface{} {
	return map[string]interface{}{"pid": obj.Pid}
}

// Redis Key .
func (obj *Post) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Pid)
}

// GetChanges .
func (obj *Post) GetChanges() map[string]interface{} {
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
func (obj *Post) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
