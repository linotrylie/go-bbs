package transform

import (
	"fmt"
	"go-bbs/app/http/model"
	"go-bbs/app/http/model/response"
	"go-bbs/global"
	"go-bbs/utils"
	"time"
)

func TransformUser(user *model.User) (userVo response.UserVo) {
	userVo.VipEnd = user.VipEnd
	userVo.Username = user.Username
	userVo.Uid = user.Uid
	userVo.LoginDate = time.Unix(int64(user.LoginDate), 0).Format(time.DateTime)
	userVo.Logins = user.Logins
	userVo.LoginIp = utils.Long2ip(user.LoginIp)
	var avatar string
	if global.CONFIG.System.OssType != "local" {
		mid := fmt.Sprintf("%09d", user.Uid)
		avatar = global.CONFIG.Local.UploadPath + "/avatar/" + mid[:3] + "/"
	} else {
		avatar = global.CONFIG.System.Host + "/" + global.CONFIG.Local.StorePath + "/avatar/"
	}
	if user.Avatar > 0 {
		userVo.Avatar = avatar + fmt.Sprintf("%d", user.Uid) + ".png?" + fmt.Sprintf("%d", user.Avatar)
	} else {
		userVo.Avatar = avatar + "avatar.png"
	}
	userVo.Favorites = user.Favorites
	userVo.Gid = user.Gid
	userVo.Qq = user.Qq
	userVo.Credits = user.Credits
	if global.User != nil && global.User.Uid == user.Uid {
		userVo.Mobile = user.Mobile
		userVo.Rmbs = user.Rmbs
		userVo.Golds = user.Golds
		userVo.Notices = user.Notices
		userVo.UnreadNotices = user.UnreadNotices
		userVo.Invitenums = user.Invitenums
	} else {
		userVo.Mobile = utils.MaskPhone(user.Mobile)
		userVo.Rmbs = 0
		userVo.Golds = 0
		userVo.Notices = 0
		userVo.UnreadNotices = 0
		userVo.Invitenums = 0
	}
	userVo.Posts = user.Posts
	userVo.Realname = user.Realname
	userVo.Email = user.Email
	userVo.Threads = user.Threads
	userVo.CreateIp = utils.Long2ip(user.CreateIp)
	userVo.CreateDate = time.Unix(int64(user.CreateDate), 0).Format(time.DateTime)
	return
}
