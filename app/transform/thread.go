package transform

import (
	"go-bbs/app/http/model"
	"go-bbs/app/http/model/response"
	"go-bbs/utils"
	"time"
)

func TransformThread(thread *model.Thread) (threadVo *response.ThreadVo) {
	if thread == nil {
		return nil
	}
	threadVo = &response.ThreadVo{}
	threadVo.Uid = thread.Uid
	threadVo.Fid = thread.Fid
	threadVo.Firstpid = thread.Firstpid
	threadVo.Favorites = thread.Favorites
	threadVo.OfferNum = thread.OfferNum
	threadVo.OfferStatus = thread.OfferStatus
	threadVo.Userip = utils.Long2ip(thread.Userip)
	threadVo.CreateDate = time.Unix(int64(thread.CreateDate), 0).Format(time.DateTime)
	threadVo.Tid = thread.Tid
	threadVo.Top = thread.Top
	threadVo.Tagids = thread.Tagids
	threadVo.TagidsTime = time.Unix(int64(thread.TagidsTime), 0).Format(time.DateTime)
	threadVo.Digest = thread.Digest
	threadVo.ActivityId = thread.ActivityId
	threadVo.Highlight = thread.Highlight
	threadVo.Subject = thread.Subject
	threadVo.Views = thread.Views
	threadVo.Posts = thread.Posts
	threadVo.Images = thread.Images
	threadVo.Subject = thread.Subject
	threadVo.Files = thread.Files
	threadVo.Subject = thread.Subject
	threadVo.Mods = thread.Mods
	threadVo.Likes = thread.Likes
	threadVo.ContentBuy = thread.ContentBuy
	threadVo.ContentGolds = thread.ContentGolds
	threadVo.Closed = thread.Closed
	threadVo.Lastuid = thread.Lastuid
	threadVo.Deleted = thread.Deleted
	threadVo.IsVote = thread.IsVote
	threadVo.LocationTr = thread.LocationTr
	threadVo.ContentBuyType = thread.ContentBuyType
	return
}
