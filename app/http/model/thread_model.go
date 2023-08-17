package model

type Thread struct {
	changes        map[string]interface{}
	Fid            int    `gorm:"column:fid" json:"fid"`
	Tid            int    `gorm:"primaryKey;column:tid" json:"tid"`
	Top            int    `gorm:"column:top" json:"top"`
	Uid            int    `gorm:"column:uid" json:"uid"`
	Userip         int    `gorm:"column:userip" json:"userip"`
	Subject        string `gorm:"column:subject" json:"subject"`
	CreateDate     int    `gorm:"column:createdate" json:"createdate"`
	LastDate       int    `gorm:"column:lastdate" json:"lastdate"`
	Views          int    `gorm:"column:views" json:"views"`
	Posts          int    `gorm:"column:posts" json:"posts"`
	Images         int    `gorm:"column:images" json:"images"`
	Files          int    `gorm:"column:files" json:"files"`
	Mods           int    `gorm:"column:mods" json:"mods"`
	Closed         int    `gorm:"column:closed" json:"closed"`
	Firstpid       int    `gorm:"column:firstpid" json:"firstpid"`
	Lastuid        int    `gorm:"column:lastuid" json:"lastuid"`
	Lastpid        int    `gorm:"column:lastpid" json:"lastpid"`
	LocationTr     string `gorm:"column:locationtr" json:"locationtr"`
	Favorites      int    `gorm:"column:favorites" json:"favorites"` // 收藏数
	Likes          int    `gorm:"column:likes" json:"likes"`         // 点赞数
	Highlight      int    `gorm:"column:highlight" json:"highlight"`
	ContentBuy     int    `gorm:"column:contentbuy" json:"contentbuy"`
	ContentBuyType int    `gorm:"column:contentbuytype" json:"contentbuytype"`
	Digest         int    `gorm:"column:digest" json:"digest"`
	Deleted        int    `gorm:"column:deleted" json:"deleted"`
	Readp          int    `gorm:"column:readp" json:"readp"`
	OfferNum       int    `gorm:"column:offernum" json:"offernum"`
	OfferStatus    int    `gorm:"column:offerstatus" json:"offerstatus"`
	Tagids         string `gorm:"column:tagids" json:"tagids"`
	TagidsTime     int    `gorm:"column:tagidstime" json:"tagidstime"`
}

func (*Thread) TableName() string {
	return "bbs_thread"
}

// Location .
func (obj *Thread) Location() map[string]interface{} {
	return map[string]interface{}{"tid": obj.Tid}
}

// GetChanges .
func (obj *Thread) GetChanges() map[string]interface{} {
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
func (obj *Thread) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
