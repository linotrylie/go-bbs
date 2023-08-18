package model

type ThreadDigest struct {
	changes map[string]interface{}
	Fid     int `gorm:"column:fid" json:"fid"`
	Tid     int `gorm:"primaryKey;column:tid" json:"tid"`
	Uid     int `gorm:"column:uid" json:"uid"`
	Digest  int `gorm:"column:digest" json:"digest"`
}

func (*ThreadDigest) TableName() string {
	return "bbs_thread_digest"
}

// Location .
func (obj *ThreadDigest) Location() map[string]interface{} {
	return map[string]interface{}{"tid": obj.Tid}
}

// GetChanges .
func (obj *ThreadDigest) GetChanges() map[string]interface{} {
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
func (obj *ThreadDigest) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
