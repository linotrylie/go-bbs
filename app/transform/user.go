package transform

import (
	"fmt"
	"go-bbs/app/http/model"
	"go-bbs/app/http/model/response"
	"go-bbs/global"
	"go-bbs/utils"
	"time"
)

func TransformUser(user *model.User) (userVo response.User) {
	userVo.VipEnd = user.VipEnd
	userVo.Username = user.Username
	userVo.Uid = user.Uid
	userVo.UnreadNotices = user.UnreadNotices
	userVo.Invitenums = user.Invitenums
	userVo.LoginDate = time.Unix(int64(user.LoginDate), 0).Format(time.DateTime)
	userVo.Logins = user.Logins
	userVo.LoginIp = utils.Long2ip(uint32(user.LoginIp))
	avatarurl := global.CONFIG.Local.UploadPath + "/avatar/"
	mid := fmt.Sprintf("%09d", user.Uid)
	userVo.Avatar = avatarurl + mid[:3] + "/" + fmt.Sprintf("%d", user.Uid) + ".png?" + fmt.Sprintf("%d", user.Avatar)
	userVo.Credits = user.Credits
	userVo.Favorites = user.Favorites
	userVo.Gid = user.Gid
	userVo.Rmbs = user.Rmbs
	userVo.Golds = user.Golds
	userVo.Notices = user.Notices
	userVo.Qq = user.Qq
	if global.User.Uid == user.Uid {
		userVo.Mobile = user.Mobile
	} else {
		userVo.Mobile = utils.MaskPhone(user.Mobile)
	}
	userVo.Posts = user.Posts
	userVo.Realname = user.Realname
	userVo.Email = user.Email
	userVo.Threads = user.Threads
	userVo.CreateIp = utils.Long2ip(uint32(user.CreateIp))
	userVo.CreateDate = time.Unix(int64(user.CreateDate), 0).Format(time.DateTime)
	return
}
