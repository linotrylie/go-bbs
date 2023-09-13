package response

type User struct {
	Uid           int    `gorm:"primaryKey;column:uid" json:"uid"`      // 用户编号
	Gid           int    `gorm:"column:gid" json:"gid"`                 // 用户组编号
	Email         string `gorm:"column:email" json:"email"`             // 邮箱
	Username      string `gorm:"column:username" json:"username"`       // 用户名
	Realname      string `gorm:"column:realname" json:"realname"`       // 用户名
	Mobile        string `gorm:"column:mobile" json:"mobile"`           // 手机号
	Qq            string `gorm:"column:qq" json:"qq"`                   // QQ
	Threads       int    `gorm:"column:threads" json:"threads"`         // 发帖数
	Posts         int    `gorm:"column:posts" json:"posts"`             // 回帖数
	Credits       int    `gorm:"column:credits" json:"credits"`         // 积分
	Golds         int    `gorm:"column:golds" json:"golds"`             // 金币
	Rmbs          int    `gorm:"column:rmbs" json:"rmbs"`               // 人民币
	CreateIp      string `gorm:"column:create_ip" json:"create_ip"`     // 创建时IP
	CreateDate    string `gorm:"column:create_date" json:"create_date"` // 创建时间
	LoginIp       string `gorm:"column:login_ip" json:"login_ip"`       // 登录时IP
	LoginDate     string `gorm:"column:login_date" json:"login_date"`   // 登录时间
	Logins        int    `gorm:"column:logins" json:"logins"`           // 登录次数
	Avatar        string `gorm:"column:avatar" json:"avatar"`           // 用户最后更新图像时间
	Invitenums    int    `gorm:"column:invitenums" json:"invitenums"`
	Favorites     int    `gorm:"column:favorites" json:"favorites"` // 收藏数
	Notices       int    `gorm:"column:notices" json:"notices"`
	UnreadNotices int    `gorm:"column:unread_notices" json:"unread_notices"`
	VipEnd        int    `gorm:"column:vip_end" json:"vip_end"`
}
