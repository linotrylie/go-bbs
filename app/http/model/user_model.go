package model

type User struct {
	changes       map[string]interface{}
	Uid           int    `gorm:"primaryKey;column:uid" json:"uid"`      // 用户编号
	Gid           int    `gorm:"column:gid" json:"gid"`                 // 用户组编号
	Email         string `gorm:"column:email" json:"email"`             // 邮箱
	Username      string `gorm:"column:username" json:"username"`       // 用户名
	Realname      string `gorm:"column:realname" json:"realname"`       // 用户名
	Password      string `gorm:"column:password" json:"password"`       // 密码
	PasswordSms   string `gorm:"column:passwordsms" json:"passwordsms"` // 密码
	Salt          string `gorm:"column:salt" json:"salt"`               // 密码混杂
	Mobile        string `gorm:"column:mobile" json:"mobile"`           // 手机号
	Qq            string `gorm:"column:qq" json:"qq"`                   // QQ
	Threads       int    `gorm:"column:threads" json:"threads"`         // 发帖数
	Posts         int    `gorm:"column:posts" json:"posts"`             // 回帖数
	Credits       int    `gorm:"column:credits" json:"credits"`         // 积分
	Golds         int    `gorm:"column:golds" json:"golds"`             // 金币
	Rmbs          int    `gorm:"column:rmbs" json:"rmbs"`               // 人民币
	CreateIp      int    `gorm:"column:createip" json:"createip"`       // 创建时IP
	CreateDate    int    `gorm:"column:createdate" json:"createdate"`   // 创建时间
	LoginIp       int    `gorm:"column:loginip" json:"loginip"`         // 登录时IP
	LoginDate     int    `gorm:"column:logindate" json:"logindate"`     // 登录时间
	Logins        int    `gorm:"column:logins" json:"logins"`           // 登录次数
	Avatar        int    `gorm:"column:avatar" json:"avatar"`           // 用户最后更新图像时间
	Invitenums    int    `gorm:"column:invitenums" json:"invitenums"`
	Favorites     int    `gorm:"column:favorites" json:"favorites"` // 收藏数
	Notices       int    `gorm:"column:notices" json:"notices"`
	UnreadNotices int    `gorm:"column:unreadnotices" json:"unreadnotices"`
	VipEnd        int    `gorm:"column:vipend" json:"vipend"`
	EmailV        string `gorm:"column:emailv" json:"emailv"`
	Digests       int    `gorm:"column:digests" json:"digests"`
	Digests3      int    `gorm:"column:digests3" json:"digests3"`
}

func (*User) TableName() string {
	return "user"
}

// Location .
func (obj *User) Location() map[string]interface{} {
	return map[string]interface{}{"Uid": obj.Uid}
}

// GetChanges .
func (obj *User) GetChanges() map[string]interface{} {
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
func (obj *User) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}
