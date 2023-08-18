package requests

type UserRequest struct {
	Uid           int    ` json:"uid"`         // 用户编号
	Gid           int    ` json:"gid"`         // 用户组编号
	Email         string ` json:"email"`       // 邮箱
	Username      string ` json:"username"`    // 用户名
	Realname      string ` json:"realname"`    // 用户名
	Password      string ` json:"password"`    // 密码
	PasswordSms   string ` json:"passwordsms"` // 密码
	Salt          string ` json:"salt"`        // 密码混杂
	Mobile        string ` json:"mobile"`      // 手机号
	Qq            string ` json:"qq"`          // QQ
	Threads       int    ` json:"threads"`     // 发帖数
	Posts         int    ` json:"posts"`       // 回帖数
	Credits       int    ` json:"credits"`     // 积分
	Golds         int    ` json:"golds"`       // 金币
	Rmbs          int    ` json:"rmbs"`        // 人民币
	CreateIp      int    ` json:"createip"`    // 创建时IP
	CreateDate    int    ` json:"createdate"`  // 创建时间
	LoginIp       int    ` json:"loginip"`     // 登录时IP
	LoginDate     int    ` json:"logindate"`   // 登录时间
	Logins        int    ` json:"logins"`      // 登录次数
	Avatar        int    ` json:"avatar"`      // 用户最后更新图像时间
	Invitenums    int    ` json:"invitenums"`
	Favorites     int    ` json:"favorites"` // 收藏数
	Notices       int    ` json:"notices"`
	UnreadNotices int    ` json:"unreadnotices"`
	VipEnd        int    ` json:"vipend"`
	EmailV        string ` json:"emailv"`
	Digests       int    ` json:"digests"`
	Digests3      int    ` json:"digests3"`
}
