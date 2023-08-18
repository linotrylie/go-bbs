package model

type Post struct {
	changes          map[string]interface{}
	Tid              int    `gorm:"column:tid" json:"tid"`
	Pid              int    `gorm:"primaryKey;column:pid" json:"pid"`
	Uid              int    `gorm:"column:uid" json:"uid"`
	Isfirst          int    `gorm:"column:isfirst" json:"isfirst"`
	CreateDate       int    `gorm:"column:createdate" json:"createdate"`
	Userip           int    `gorm:"column:userip" json:"userip"`
	Images           int    `gorm:"column:images" json:"images"`
	Files            int    `gorm:"column:files" json:"files"`
	Doctype          int    `gorm:"column:doctype" json:"doctype"`
	Quotepid         int    `gorm:"column:quotepid" json:"quotepid"`
	Message          string `gorm:"column:message" json:"message"`
	MessageFmt       string `gorm:"column:messagefmt" json:"messagefmt"`
	LocationPost     string `gorm:"column:locationpost" json:"locationpost"`
	Likes            int    `gorm:"column:likes" json:"likes"` // 点赞数
	Deleted          int    `gorm:"column:deleted" json:"deleted"`
	Updates          int    `gorm:"column:updates" json:"updates"`
	LastUpdateDate   int    `gorm:"column:lastupdatedate" json:"lastupdatedate"`
	LastUpdateUid    int    `gorm:"column:lastupdateuid" json:"lastupdateuid"`
	LastUpdateReason string `gorm:"column:lastupdatereason" json:"lastupdatereason"`
}

func (*Post) TableName() string {
	return "bbs_post"
}

// Location .
func (obj *Post) Location() map[string]interface{} {
	return map[string]interface{}{"pid": obj.Pid}
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
