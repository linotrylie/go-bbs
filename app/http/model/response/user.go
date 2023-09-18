package response

import "go-bbs/app/http/model"

type UserVo struct {
	Uid           int             `json:"uid"`         // 用户编号
	Gid           int             `json:"gid"`         // 用户组编号
	Email         string          `json:"email"`       // 邮箱
	Username      string          `json:"username"`    // 用户名
	Realname      string          `json:"realname"`    // 用户名
	Mobile        string          `json:"mobile"`      // 手机号
	Qq            string          `json:"qq"`          // QQ
	Threads       int             `json:"threads"`     // 发帖数
	Posts         int             `json:"posts"`       // 回帖数
	Credits       int             `json:"credits"`     // 积分
	Golds         int             `json:"golds"`       // 金币
	Rmbs          int             `json:"rmbs"`        // 人民币
	CreateIp      string          `json:"create_ip"`   // 创建时IP
	CreateDate    string          `json:"create_date"` // 创建时间
	LoginIp       string          `json:"login_ip"`    // 登录时IP
	LoginDate     string          `json:"login_date"`  // 登录时间
	Logins        int             `json:"logins"`      // 登录次数
	Avatar        string          `json:"avatar"`      // 用户最后更新图像时间
	Invitenums    int             `json:"invitenums"`
	Favorites     int             `json:"favorites"` // 收藏数
	Notices       int             `json:"notices"`
	UnreadNotices int             `json:"unread_notices"`
	VipEnd        int             `json:"vip_end"`
	Group         *model.Group    `json:"group,omitempty"`
	ThreadList    []*model.Thread `json:"thread_list,omitempty"`
}
