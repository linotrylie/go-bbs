package model

import (
	"fmt"
)

type Post struct {
	changes          map[string]interface{}
	Tid              int    `gorm:"column:tid" json:"tid"`
	Pid              int    `gorm:"primaryKey;column:pid" json:"pid"`
	Uid              int    `gorm:"column:uid" json:"uid"`
	Isfirst          int    `gorm:"column:isfirst" json:"isfirst"`
	CreateDate       int64  `gorm:"column:create_date" json:"create_date"`
	Userip           uint32 `gorm:"column:userip" json:"userip"`
	Images           int    `gorm:"column:images" json:"images"`
	Files            int    `gorm:"column:files" json:"files"`
	Doctype          int    `gorm:"column:doctype" json:"doctype"`
	Quotepid         int    `gorm:"column:quotepid" json:"quotepid"`
	Message          string `gorm:"column:message" json:"message"`
	MessageFmt       string `gorm:"column:message_fmt" json:"message_fmt"`
	LocationPost     string `gorm:"column:location_post" json:"location_post"`
	Likes            int    `gorm:"column:likes" json:"likes"` // 点赞数
	Deleted          int    `gorm:"column:deleted" json:"deleted"`
	Updates          int    `gorm:"column:updates" json:"updates"`
	LastUpdateDate   int64  `gorm:"column:last_update_date" json:"last_update_date"`
	LastUpdateUid    int    `gorm:"column:last_update_uid" json:"last_update_uid"`
	LastUpdateReason string `gorm:"column:last_update_reason" json:"last_update_reason"`
	ReplyHide        int    `gorm:"column:reply_hide" json:"reply_hide"`
	LastUpdateUser   *User  `gorm:"foreignkey:last_update_uid;references:uid"`
	CreateUser       *User  `gorm:"foreignkey:uid;references:uid"`
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
func (obj *Post) SetTid(val int) *Post {
	obj.Tid = val
	obj.Update("tid", obj.Tid)
	return obj
}
func (obj *Post) SetPid(val int) *Post {
	obj.Pid = val
	obj.Update("pid", obj.Pid)
	return obj
}
func (obj *Post) SetUid(val int) *Post {
	obj.Uid = val
	obj.Update("uid", obj.Uid)
	return obj
}
func (obj *Post) SetIsfirst(val int) *Post {
	obj.Isfirst += val
	obj.Update("isfirst", obj.Isfirst)
	return obj
}
func (obj *Post) SetCreateDate(val int64) *Post {
	obj.CreateDate += val
	obj.Update("create_date", obj.CreateDate)
	return obj
}
func (obj *Post) SetUserip(val uint32) *Post {
	obj.Userip += val
	obj.Update("userip", obj.Userip)
	return obj
}
func (obj *Post) SetImages(val int) *Post {
	obj.Images += val
	obj.Update("images", obj.Images)
	return obj
}
func (obj *Post) SetFiles(val int) *Post {
	obj.Files += val
	obj.Update("files", obj.Files)
	return obj
}
func (obj *Post) SetDoctype(val int) *Post {
	obj.Doctype += val
	obj.Update("doctype", obj.Doctype)
	return obj
}
func (obj *Post) SetQuotepid(val int) *Post {
	obj.Quotepid = val
	obj.Update("quotepid", obj.Quotepid)
	return obj
}
func (obj *Post) SetMessage(val string) *Post {
	obj.Message = val
	obj.Update("message", obj.Message)
	return obj
}
func (obj *Post) SetMessageFmt(val string) *Post {
	obj.MessageFmt = val
	obj.Update("message_fmt", obj.MessageFmt)
	return obj
}
func (obj *Post) SetLocationPost(val string) *Post {
	obj.LocationPost = val
	obj.Update("location_post", obj.LocationPost)
	return obj
}
func (obj *Post) SetLikes(val int) *Post {
	obj.Likes += val
	obj.Update("likes", obj.Likes)
	return obj
}
func (obj *Post) SetDeleted(val int) *Post {
	obj.Deleted += val
	obj.Update("deleted", obj.Deleted)
	return obj
}
func (obj *Post) SetUpdates(val int) *Post {
	obj.Updates += val
	obj.Update("updates", obj.Updates)
	return obj
}
func (obj *Post) SetLastUpdateDate(val int64) *Post {
	obj.LastUpdateDate += val
	obj.Update("last_update_date", obj.LastUpdateDate)
	return obj
}
func (obj *Post) SetLastUpdateUid(val int) *Post {
	obj.LastUpdateUid = val
	obj.Update("last_update_uid", obj.LastUpdateUid)
	return obj
}
func (obj *Post) SetLastUpdateReason(val string) *Post {
	obj.LastUpdateReason = val
	obj.Update("last_update_reason", obj.LastUpdateReason)
	return obj
}
func (obj *Post) SetReplyHide(val int) *Post {
	obj.ReplyHide = val
	obj.Update("reply_hide", obj.ReplyHide)
	return obj
}
