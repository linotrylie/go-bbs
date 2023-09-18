package model

import (
	"fmt"
)

type Attach struct {
	changes     map[string]interface{}
	Aid         int    `gorm:"primaryKey;column:aid" json:"aid"`
	Tid         int    `gorm:"column:tid" json:"tid"`
	Pid         int    `gorm:"column:pid" json:"pid"`
	Uid         int    `gorm:"column:uid" json:"uid"`
	Filesize    int64  `gorm:"column:filesize" json:"filesize"`
	Width       int    `gorm:"column:width" json:"width"`
	Height      int    `gorm:"column:height" json:"height"`
	Filename    string `gorm:"column:filename" json:"filename"`
	Orgfilename string `gorm:"column:orgfilename" json:"orgfilename"`
	Filetype    string `gorm:"column:filetype" json:"filetype"`
	CreateDate  int64  `gorm:"column:create_date" json:"create_date"`
	Comment     string `gorm:"column:comment" json:"comment"`
	Downloads   int    `gorm:"column:downloads" json:"downloads"`
	Credits     int    `gorm:"column:credits" json:"credits"`
	Golds       int    `gorm:"column:golds" json:"golds"`
	Rmbs        int    `gorm:"column:rmbs" json:"rmbs"`
	Isimage     int    `gorm:"column:isimage" json:"isimage"`
}

func (*Attach) TableName() string {
	return "bbs_attach"
}

// Location .
func (obj *Attach) Location() map[string]interface{} {
	return map[string]interface{}{"aid": obj.Aid}
}

// Redis Key .
func (obj *Attach) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Aid)
}

// GetChanges .
func (obj *Attach) GetChanges() map[string]interface{} {
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
func (obj *Attach) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
func (obj *Attach) SetAid(val int) *Attach {
	obj.Aid = val
	obj.Update("aid", obj.Aid)
	return obj
}
func (obj *Attach) SetTid(val int) *Attach {
	obj.Tid = val
	obj.Update("tid", obj.Tid)
	return obj
}
func (obj *Attach) SetPid(val int) *Attach {
	obj.Pid = val
	obj.Update("pid", obj.Pid)
	return obj
}
func (obj *Attach) SetUid(val int) *Attach {
	obj.Uid = val
	obj.Update("uid", obj.Uid)
	return obj
}
func (obj *Attach) SetFilesize(val int64) *Attach {
	obj.Filesize = val
	obj.Update("filesize", obj.Filesize)
	return obj
}
func (obj *Attach) SetWidth(val int) *Attach {
	obj.Width = val
	obj.Update("width", obj.Width)
	return obj
}
func (obj *Attach) SetHeight(val int) *Attach {
	obj.Height += val
	obj.Update("height", obj.Height)
	return obj
}
func (obj *Attach) SetFilename(val string) *Attach {
	obj.Filename = val
	obj.Update("filename", obj.Filename)
	return obj
}
func (obj *Attach) SetOrgfilename(val string) *Attach {
	obj.Orgfilename = val
	obj.Update("orgfilename", obj.Orgfilename)
	return obj
}
func (obj *Attach) SetFiletype(val string) *Attach {
	obj.Filetype = val
	obj.Update("filetype", obj.Filetype)
	return obj
}
func (obj *Attach) SetCreateDate(val int64) *Attach {
	obj.CreateDate += val
	obj.Update("create_date", obj.CreateDate)
	return obj
}
func (obj *Attach) SetComment(val string) *Attach {
	obj.Comment = val
	obj.Update("comment", obj.Comment)
	return obj
}
func (obj *Attach) SetDownloads(val int) *Attach {
	obj.Downloads += val
	obj.Update("downloads", obj.Downloads)
	return obj
}
func (obj *Attach) SetCredits(val int) *Attach {
	obj.Credits += val
	obj.Update("credits", obj.Credits)
	return obj
}
func (obj *Attach) SetGolds(val int) *Attach {
	obj.Golds += val
	obj.Update("golds", obj.Golds)
	return obj
}
func (obj *Attach) SetRmbs(val int) *Attach {
	obj.Rmbs += val
	obj.Update("rmbs", obj.Rmbs)
	return obj
}
func (obj *Attach) SetIsimage(val int) *Attach {
	obj.Isimage += val
	obj.Update("isimage", obj.Isimage)
	return obj
}
