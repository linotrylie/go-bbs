package model

import (
	"fmt"
)

type Thread struct {
	changes        map[string]interface{}
	Fid            int    `gorm:"column:fid" json:"fid"`
	Tid            int    `gorm:"primaryKey;column:tid" json:"tid"`
	Top            int    `gorm:"column:top" json:"top"`
	Uid            int    `gorm:"column:uid" json:"uid"`
	Userip         uint32 `gorm:"column:userip" json:"userip"`
	Subject        string `gorm:"column:subject" json:"subject"`
	CreateDate     int64  `gorm:"column:create_date" json:"create_date"`
	LastDate       int64  `gorm:"column:last_date" json:"last_date"`
	Views          int    `gorm:"column:views" json:"views"`
	Posts          int    `gorm:"column:posts" json:"posts"`
	Images         int    `gorm:"column:images" json:"images"`
	Files          int    `gorm:"column:files" json:"files"`
	Mods           int    `gorm:"column:mods" json:"mods"`
	Closed         int    `gorm:"column:closed" json:"closed"`
	Firstpid       int    `gorm:"column:firstpid" json:"firstpid"`
	Lastuid        int    `gorm:"column:lastuid" json:"lastuid"`
	Lastpid        int    `gorm:"column:lastpid" json:"lastpid"`
	LocationTr     string `gorm:"column:location_tr" json:"location_tr"`
	Favorites      int    `gorm:"column:favorites" json:"favorites"` // 收藏数
	Likes          int    `gorm:"column:likes" json:"likes"`         // 点赞数
	Highlight      int    `gorm:"column:highlight" json:"highlight"`
	ContentBuy     int    `gorm:"column:content_buy" json:"content_buy"`
	ContentBuyType int    `gorm:"column:content_buy_type" json:"content_buy_type"`
	Digest         int    `gorm:"column:digest" json:"digest"`
	Deleted        int    `gorm:"column:deleted" json:"deleted"`
	Readp          int    `gorm:"column:readp" json:"readp"`
	OfferNum       int    `gorm:"column:offernum" json:"offernum"`
	OfferStatus    int    `gorm:"column:offerstatus" json:"offerstatus"`
	Tagids         string `gorm:"column:tagids" json:"tagids"`
	TagidsTime     int64  `gorm:"column:tagids_time" json:"tagids_time"`
	IsVote         int    `gorm:"column:is_vote" json:"is_vote"`
	AttachGolds    int    `gorm:"column:attach_golds" json:"attach_golds"`
	ContentGolds   int    `gorm:"column:content_golds" json:"content_golds"`
	User           User   `gorm:"foreignkey:uid;references:uid"`
}

func (*Thread) TableName() string {
	return "bbs_thread"
}

// Location .
func (obj *Thread) Location() map[string]interface{} {
	return map[string]interface{}{"tid": obj.Tid}
}

// Redis Key .
func (obj *Thread) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Tid)
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
func (obj *Thread) SetFid(val int) *Thread {
	obj.Fid = val
	obj.Update("fid", obj.Fid)
	return obj
}
func (obj *Thread) SetTid(val int) *Thread {
	obj.Tid = val
	obj.Update("tid", obj.Tid)
	return obj
}
func (obj *Thread) SetTop(val int) *Thread {
	obj.Top += val
	obj.Update("top", obj.Top)
	return obj
}
func (obj *Thread) SetUid(val int) *Thread {
	obj.Uid = val
	obj.Update("uid", obj.Uid)
	return obj
}
func (obj *Thread) SetUserip(val uint32) *Thread {
	obj.Userip += val
	obj.Update("userip", obj.Userip)
	return obj
}
func (obj *Thread) SetSubject(val string) *Thread {
	obj.Subject = val
	obj.Update("subject", obj.Subject)
	return obj
}
func (obj *Thread) SetCreateDate(val int64) *Thread {
	obj.CreateDate += val
	obj.Update("create_date", obj.CreateDate)
	return obj
}
func (obj *Thread) SetLastDate(val int64) *Thread {
	obj.LastDate += val
	obj.Update("last_date", obj.LastDate)
	return obj
}
func (obj *Thread) SetViews(val int) *Thread {
	obj.Views += val
	obj.Update("views", obj.Views)
	return obj
}
func (obj *Thread) SetPosts(val int) *Thread {
	obj.Posts += val
	obj.Update("posts", obj.Posts)
	return obj
}
func (obj *Thread) SetImages(val int) *Thread {
	obj.Images += val
	obj.Update("images", obj.Images)
	return obj
}
func (obj *Thread) SetFiles(val int) *Thread {
	obj.Files += val
	obj.Update("files", obj.Files)
	return obj
}
func (obj *Thread) SetMods(val int) *Thread {
	obj.Mods += val
	obj.Update("mods", obj.Mods)
	return obj
}
func (obj *Thread) SetClosed(val int) *Thread {
	obj.Closed += val
	obj.Update("closed", obj.Closed)
	return obj
}
func (obj *Thread) SetFirstpid(val int) *Thread {
	obj.Firstpid = val
	obj.Update("firstpid", obj.Firstpid)
	return obj
}
func (obj *Thread) SetLastuid(val int) *Thread {
	obj.Lastuid = val
	obj.Update("lastuid", obj.Lastuid)
	return obj
}
func (obj *Thread) SetLastpid(val int) *Thread {
	obj.Lastpid = val
	obj.Update("lastpid", obj.Lastpid)
	return obj
}
func (obj *Thread) SetLocationTr(val string) *Thread {
	obj.LocationTr = val
	obj.Update("location_tr", obj.LocationTr)
	return obj
}
func (obj *Thread) SetFavorites(val int) *Thread {
	obj.Favorites += val
	obj.Update("favorites", obj.Favorites)
	return obj
}
func (obj *Thread) SetLikes(val int) *Thread {
	obj.Likes += val
	obj.Update("likes", obj.Likes)
	return obj
}
func (obj *Thread) SetHighlight(val int) *Thread {
	obj.Highlight += val
	obj.Update("highlight", obj.Highlight)
	return obj
}
func (obj *Thread) SetContentBuy(val int) *Thread {
	obj.ContentBuy += val
	obj.Update("content_buy", obj.ContentBuy)
	return obj
}
func (obj *Thread) SetContentBuyType(val int) *Thread {
	obj.ContentBuyType += val
	obj.Update("content_buy_type", obj.ContentBuyType)
	return obj
}
func (obj *Thread) SetDigest(val int) *Thread {
	obj.Digest += val
	obj.Update("digest", obj.Digest)
	return obj
}
func (obj *Thread) SetDeleted(val int) *Thread {
	obj.Deleted += val
	obj.Update("deleted", obj.Deleted)
	return obj
}
func (obj *Thread) SetReadp(val int) *Thread {
	obj.Readp += val
	obj.Update("readp", obj.Readp)
	return obj
}
func (obj *Thread) SetOfferNum(val int) *Thread {
	obj.OfferNum += val
	obj.Update("offernum", obj.OfferNum)
	return obj
}
func (obj *Thread) SetOfferStatus(val int) *Thread {
	obj.OfferStatus += val
	obj.Update("offerstatus", obj.OfferStatus)
	return obj
}
func (obj *Thread) SetTagids(val string) *Thread {
	obj.Tagids = val
	obj.Update("tagids", obj.Tagids)
	return obj
}
func (obj *Thread) SetTagidsTime(val int64) *Thread {
	obj.TagidsTime = val
	obj.Update("tagids_time", obj.TagidsTime)
	return obj
}
func (obj *Thread) SetIsVote(val int) *Thread {
	obj.IsVote += val
	obj.Update("is_vote", obj.IsVote)
	return obj
}
func (obj *Thread) SetAttachGolds(val int) *Thread {
	obj.AttachGolds += val
	obj.Update("attach_golds", obj.AttachGolds)
	return obj
}
func (obj *Thread) SetContentGolds(val int) *Thread {
	obj.ContentGolds += val
	obj.Update("content_golds", obj.ContentGolds)
	return obj
}
