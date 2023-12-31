package model

import (
	"fmt"
)

type Kami struct {
	changes map[string]interface{}
	Kahao   int    `gorm:"primaryKey;column:kahao" json:"kahao"`
	Kami    string `gorm:"column:kami" json:"kami"`
	Mianzhi int    `gorm:"column:mianzhi" json:"mianzhi"`
	Uid     int    `gorm:"column:uid" json:"uid"`
	Riqi    string `gorm:"column:riqi" json:"riqi"`
}

func (*Kami) TableName() string {
	return "bbs_kami"
}

// Location .
func (obj *Kami) Location() map[string]interface{} {
	return map[string]interface{}{"kahao": obj.Kahao}
}

// Redis Key .
func (obj *Kami) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Kahao)
}

// GetChanges .
func (obj *Kami) GetChanges() map[string]interface{} {
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
func (obj *Kami) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
func (obj *Kami) SetKahao(val int) *Kami {
	obj.Kahao += val
	obj.Update("kahao", obj.Kahao)
	return obj
}
func (obj *Kami) SetKami(val string) *Kami {
	obj.Kami = val
	obj.Update("kami", obj.Kami)
	return obj
}
func (obj *Kami) SetMianzhi(val int) *Kami {
	obj.Mianzhi += val
	obj.Update("mianzhi", obj.Mianzhi)
	return obj
}
func (obj *Kami) SetUid(val int) *Kami {
	obj.Uid = val
	obj.Update("uid", obj.Uid)
	return obj
}
func (obj *Kami) SetRiqi(val string) *Kami {
	obj.Riqi = val
	obj.Update("riqi", obj.Riqi)
	return obj
}
