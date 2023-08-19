package model

type PostLike struct {
	changes    map[string]interface{}
	Tid        int `gorm:"column:tid" json:"tid"`               // 帖子ID
	Pid        int `gorm:"column:pid" json:"pid"`               // 回帖ID
	Uid        int `gorm:"column:uid" json:"uid"`               // 用户ID
	CreateDate int `gorm:"column:createdate" json:"createdate"` // 添加时间
	CreateIp   int `gorm:"column:createip" json:"createip"`     // 添加IP
}

func (*PostLike) TableName() string {
	return "bbs_post_like"
}

// Location .
func (obj *PostLike) Location() map[string]interface{} {
	return map[string]interface{}{}
}

// GetChanges .
func (obj *PostLike) GetChanges() map[string]interface{} {
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
func (obj *PostLike) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
