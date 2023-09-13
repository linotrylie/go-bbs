package model

import (
	"fmt"
)

type User struct {
	changes       map[string]interface{}
	Uid           int    `gorm:"primaryKey;column:uid" json:"uid"`        // 用户编号
	Gid           int    `gorm:"column:gid" json:"gid"`                   // 用户组编号
	Email         string `gorm:"column:email" json:"email"`               // 邮箱
	Username      string `gorm:"column:username" json:"username"`         // 用户名
	Realname      string `gorm:"column:realname" json:"realname"`         // 用户名
	Password      string `gorm:"column:password" json:"password"`         // 密码
	PasswordSms   string `gorm:"column:password_sms" json:"password_sms"` // 密码
	Salt          string `gorm:"column:salt" json:"salt"`                 // 密码混杂
	Mobile        string `gorm:"column:mobile" json:"mobile"`             // 手机号
	Qq            string `gorm:"column:qq" json:"qq"`                     // QQ
	Threads       int    `gorm:"column:threads" json:"threads"`           // 发帖数
	Posts         int    `gorm:"column:posts" json:"posts"`               // 回帖数
	Credits       int    `gorm:"column:credits" json:"credits"`           // 积分
	Golds         int    `gorm:"column:golds" json:"golds"`               // 金币
	Rmbs          int    `gorm:"column:rmbs" json:"rmbs"`                 // 人民币
	CreateIp      int    `gorm:"column:create_ip" json:"create_ip"`       // 创建时IP
	CreateDate    int    `gorm:"column:create_date" json:"create_date"`   // 创建时间
	LoginIp       int    `gorm:"column:login_ip" json:"login_ip"`         // 登录时IP
	LoginDate     int    `gorm:"column:login_date" json:"login_date"`     // 登录时间
	Logins        int    `gorm:"column:logins" json:"logins"`             // 登录次数
	Avatar        int    `gorm:"column:avatar" json:"avatar"`             // 用户最后更新图像时间
	Invitenums    int    `gorm:"column:invitenums" json:"invitenums"`
	Favorites     int    `gorm:"column:favorites" json:"favorites"` // 收藏数
	Notices       int    `gorm:"column:notices" json:"notices"`
	UnreadNotices int    `gorm:"column:unread_notices" json:"unread_notices"`
	VipEnd        int    `gorm:"column:vip_end" json:"vip_end"`
	EmailV        string `gorm:"column:email_v" json:"email_v"`
	Digests       int    `gorm:"column:digests" json:"digests"`
	Digests3      int    `gorm:"column:digests_3" json:"digests_3"`
	Signature     string `gorm:"column:signature" json:"signature"` // 用户签名
}

func (*User) TableName() string {
	return "bbs_user"
}

// Location .
func (obj *User) Location() map[string]interface{} {
	return map[string]interface{}{"uid": obj.Uid}
}

// RedisKey .
func (obj *User) RedisKey() string {
	return obj.TableName() + "_" + fmt.Sprintf("%v", obj.Uid)
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

// SetCredits .
func (obj *User) SetCredits(n int) *User {
	obj.Credits += n
	obj.Update("credits", obj.Credits)
	return obj
}

// SetGolds .
func (obj *User) SetGolds(n int) *User {
	obj.Golds += n
	obj.Update("golds", obj.Golds)
	return obj
}

// SetRmbs .
func (obj *User) SetRmbs(n int) *User {
	obj.Rmbs += n
	obj.Update("rmbs", obj.Rmbs)
	return obj
}

func (obj *User) SetPosts(n int) *User {
	obj.Posts += n
	obj.Update("posts", obj.Posts)
	return obj
}

func (obj *User) SetThreads(n int) *User {
	obj.Threads += n
	obj.Update("threads", obj.Threads)
	return obj
}
func (obj *User) SetLogins(n int) *User {
	obj.Logins += n
	obj.Update("logins", obj.Logins)
	return obj
}

func (obj *User) SetLoginIP(ip int) *User {
	obj.LoginIp = ip
	obj.Update("login_ip", obj.LoginIp)
	return obj
}
func (obj *User) SetLoginDate(time int) *User {
	obj.LoginDate = time
	obj.Update("login_date", obj.LoginDate)
	return obj
}
