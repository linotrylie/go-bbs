package model

type SgSignSet struct {
	changes   map[string]interface{}
	Id        int    `gorm:"primaryKey;column:id" json:"id"`    // id
	SgSignnum int    `gorm:"column:sgsignnum" json:"sgsignnum"` // 签到总数
	SgSign    int    `gorm:"column:sgsign" json:"sgsign"`       // 今日签到人数
	SgSignOne string `gorm:"column:sgsignone" json:"sgsignone"` // 今日第一
	SgSignTop string `gorm:"column:sgsigntop" json:"sgsigntop"` // 今日前十
	Time      int    `gorm:"column:time" json:"time"`           // 最后签到时间
}

func (*SgSignSet) TableName() string {
	return "bbs_sg_sign_set"
}

// Location .
func (obj *SgSignSet) Location() map[string]interface{} {
	return map[string]interface{}{"id": obj.Id}
}

// GetChanges .
func (obj *SgSignSet) GetChanges() map[string]interface{} {
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
func (obj *SgSignSet) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
