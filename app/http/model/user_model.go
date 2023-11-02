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
	CreateIp      uint32 `gorm:"column:create_ip" json:"create_ip"`       // 创建时IP
	CreateDate    int64  `gorm:"column:create_date" json:"create_date"`   // 创建时间
	LoginIp       uint32 `gorm:"column:login_ip" json:"login_ip"`         // 登录时IP
	LoginDate     int64  `gorm:"column:login_date" json:"login_date"`     // 登录时间
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
	Signature     string `gorm:"column:signature" json:"signature"`       // 用户签名
	MachineCode   string `gorm:"column:machine_code" json:"machine_code"` // 机器码
	KadaoTime     int64  `gorm:"column:kadao_time" json:"kadao_time"`     // 卡刀时间
}

func (*User) TableName() string {
	return "bbs_user"
}

// Location .
func (obj *User) Location() map[string]interface{} {
	return map[string]interface{}{"uid": obj.Uid}
}

// Redis Key .
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
func (obj *User) SetUid(val int) *User {
	obj.Uid = val
	obj.Update("uid", obj.Uid)
	return obj
}
func (obj *User) SetGid(val int) *User {
	obj.Gid = val
	obj.Update("gid", obj.Gid)
	return obj
}
func (obj *User) SetEmail(val string) *User {
	obj.Email = val
	obj.Update("email", obj.Email)
	return obj
}
func (obj *User) SetUsername(val string) *User {
	obj.Username = val
	obj.Update("username", obj.Username)
	return obj
}
func (obj *User) SetRealname(val string) *User {
	obj.Realname = val
	obj.Update("realname", obj.Realname)
	return obj
}
func (obj *User) SetPassword(val string) *User {
	obj.Password = val
	obj.Update("password", obj.Password)
	return obj
}
func (obj *User) SetPasswordSms(val string) *User {
	obj.PasswordSms = val
	obj.Update("password_sms", obj.PasswordSms)
	return obj
}
func (obj *User) SetSalt(val string) *User {
	obj.Salt = val
	obj.Update("salt", obj.Salt)
	return obj
}
func (obj *User) SetMobile(val string) *User {
	obj.Mobile = val
	obj.Update("mobile", obj.Mobile)
	return obj
}
func (obj *User) SetQq(val string) *User {
	obj.Qq = val
	obj.Update("qq", obj.Qq)
	return obj
}
func (obj *User) SetThreads(val int) *User {
	obj.Threads += val
	obj.Update("threads", obj.Threads)
	return obj
}
func (obj *User) SetPosts(val int) *User {
	obj.Posts += val
	obj.Update("posts", obj.Posts)
	return obj
}
func (obj *User) SetCredits(val int) *User {
	obj.Credits += val
	obj.Update("credits", obj.Credits)
	return obj
}
func (obj *User) SetGolds(val int) *User {
	obj.Golds += val
	obj.Update("golds", obj.Golds)
	return obj
}
func (obj *User) SetRmbs(val int) *User {
	obj.Rmbs += val
	obj.Update("rmbs", obj.Rmbs)
	return obj
}
func (obj *User) SetCreateIp(val uint32) *User {
	obj.CreateIp = val
	obj.Update("create_ip", obj.CreateIp)
	return obj
}
func (obj *User) SetCreateDate(val int64) *User {
	obj.CreateDate = val
	obj.Update("create_date", obj.CreateDate)
	return obj
}
func (obj *User) SetLoginIp(val uint32) *User {
	obj.LoginIp = val
	obj.Update("login_ip", obj.LoginIp)
	return obj
}
func (obj *User) SetLoginDate(val int64) *User {
	obj.LoginDate = val
	obj.Update("login_date", obj.LoginDate)
	return obj
}
func (obj *User) SetLogins(val int) *User {
	obj.Logins += val
	obj.Update("logins", obj.Logins)
	return obj
}
func (obj *User) SetAvatar(val int) *User {
	obj.Avatar += val
	obj.Update("avatar", obj.Avatar)
	return obj
}
func (obj *User) SetInvitenums(val int) *User {
	obj.Invitenums += val
	obj.Update("invitenums", obj.Invitenums)
	return obj
}
func (obj *User) SetFavorites(val int) *User {
	obj.Favorites += val
	obj.Update("favorites", obj.Favorites)
	return obj
}
func (obj *User) SetNotices(val int) *User {
	obj.Notices += val
	obj.Update("notices", obj.Notices)
	return obj
}
func (obj *User) SetUnreadNotices(val int) *User {
	obj.UnreadNotices += val
	obj.Update("unread_notices", obj.UnreadNotices)
	return obj
}
func (obj *User) SetVipEnd(val int) *User {
	obj.VipEnd += val
	obj.Update("vip_end", obj.VipEnd)
	return obj
}
func (obj *User) SetEmailV(val string) *User {
	obj.EmailV = val
	obj.Update("email_v", obj.EmailV)
	return obj
}
func (obj *User) SetDigests(val int) *User {
	obj.Digests += val
	obj.Update("digests", obj.Digests)
	return obj
}
func (obj *User) SetDigests3(val int) *User {
	obj.Digests3 += val
	obj.Update("digests_3", obj.Digests3)
	return obj
}
func (obj *User) SetSignature(val string) *User {
	obj.Signature = val
	obj.Update("signature", obj.Signature)
	return obj
}
func (obj *User) SetMachineCode(val string) *User {
	obj.MachineCode = val
	obj.Update("machine_code", obj.MachineCode)
	return obj
}
func (obj *User) SetKadaoTime(val int64) *User {
	obj.KadaoTime = val
	obj.Update("kadao_time", obj.KadaoTime)
	return obj
}
