package transform

import (
	"go-bbs/app/http/model"
	"go-bbs/app/http/model/response"
	"go-bbs/utils"
	"time"
)

func TransformPost(post *model.Post) (postVo *response.PostVo) {
	if post != nil {
		return nil
	}
	postVo = &response.PostVo{
		Tid:              post.Tid,
		Pid:              post.Pid,
		Uid:              post.Uid,
		Isfirst:          post.Isfirst,
		CreateDate:       time.Unix(post.CreateDate, 0).Format(time.DateTime),
		Userip:           utils.Long2ip(post.Userip),
		Images:           post.Images,
		Files:            post.Files,
		Doctype:          post.Doctype,
		Quotepid:         post.Quotepid,
		Message:          post.Message,
		MessageFmt:       post.MessageFmt,
		LocationPost:     post.LocationPost,
		Likes:            post.Likes,
		Deleted:          post.Deleted,
		Updates:          post.Updates,
		LastUpdateDate:   time.Unix(post.LastUpdateDate, 0).Format(time.DateTime),
		LastUpdateUid:    post.LastUpdateUid,
		LastUpdateReason: post.LastUpdateReason,
		ReplyHide:        post.ReplyHide,
	}
	return
}
