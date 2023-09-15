package service

import (
	"context"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/random"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/gin-gonic/gin"
	"go-bbs/app/exceptions"
	"go-bbs/app/http/model"
	"go-bbs/app/http/model/requests"
	"go-bbs/app/respository"
	"go-bbs/global"
	"go-bbs/utils"
	"time"
)

type UserService struct {
	UserRepo respository.UserRepository
}

func (serv *UserService) Login(user *model.User, ctx *gin.Context) (userReturn *model.User, jwtCustomClaims *JwtCustomClaims, token string, e error) {
	//先检查是否存在相同用户名的用户
	serv.UserRepo.User = &model.User{}
	where := make(map[string]interface{})
	where["username"] = user.Username
	e = serv.UserRepo.FindUserByMap(where)
	if e != nil {
		return nil, nil, "", e
	}
	if serv.UserRepo.User == nil {
		return nil, nil, "", exceptions.UserNotFound
	}
	ok := serv.VerifyPassword(serv.UserRepo.User, user.Password)
	if !ok {
		return nil, nil, "", exceptions.FailedVerify
	}
	ServiceGroupApp.JwtService.SigningKey = []byte(global.CONFIG.JWT.SigningKey)
	claims := ServiceGroupApp.JwtService.CreateClaims(serv.UserRepo.User)
	token, e = ServiceGroupApp.JwtService.CreateToken(claims)
	if e != nil {
		return nil, nil, "", e
	}
	userReturn = serv.UserRepo.User
	global.User = serv.UserRepo.User
	jwtCustomClaims = &claims
	//将用户登录信息记录在redis中
	global.REDIS.Set(context.Background(), serv.UserRepo.User.Username, "login", time.Duration(utils.DatetimeToUnix(claims.ExpiresAt.Format(time.DateTime))-time.Now().Unix())*time.Second)
	//用户通过验证后，对用户进行后续操作，如增加经验积分或者记录登录ip等等
	go serv.LoginAfter(ctx)
	return
}

func (serv *UserService) GeneratePassword(user *model.User, password string) {
	var salt string
	if strutil.IsBlank(user.Salt) {
		salt = random.RandString(16)
	} else {
		salt = user.Salt
	}
	str := password + salt
	user.SetPassword(cryptor.Md5String(str)).SetSalt(salt)
}

func (serv *UserService) VerifyPassword(user *model.User, password string) bool {
	str := password + user.Salt
	if cryptor.Md5String(str) == user.Password {
		return true
	}
	return false
}

func (serv *UserService) LoginAfter(ctx *gin.Context) {
	serv.UserRepo.User.SetLogins(1).
		SetLoginDate(int(time.Now().Unix())).
		SetLoginIP(int(utils.Ip2long(ctx.ClientIP())))
	_, err := serv.UserRepo.Update()
	if err != nil {
		return
	}
}

func (serv *UserService) Logout() {
	global.REDIS.Del(context.Background(), global.User.Username)
}

func (serv *UserService) Detail(uid int) (*model.User, error) {
	serv.UserRepo.User = &model.User{Uid: uid}
	err := serv.UserRepo.First()
	if err != nil {
		return nil, err
	}
	return serv.UserRepo.User, nil
}

func (serv *UserService) ChangesPassword(userChangePassword *requests.UserChangePassword) (err error) {
	serv.UserRepo.User = &model.User{}
	serv.UserRepo.User.Uid = global.User.Uid
	err = serv.UserRepo.First()
	if err != nil {
		return
	}
	ok := serv.VerifyPassword(serv.UserRepo.User, userChangePassword.OldPassword)
	if !ok {
		err = exceptions.FailedVerify
		return
	}
	serv.GeneratePassword(serv.UserRepo.User, userChangePassword.NewPassword)
	update, err := serv.UserRepo.Update()
	if err != nil {
		return err
	}
	if update < 1 {
		err = exceptions.ModifyError
		return
	}
	//修改密码后，需要重新登录
	serv.Logout()
	return nil
}

func (serv *UserService) Edit(userEdit *requests.UserEdit) (err error) {

	serv.UserRepo.User = &model.User{Uid: global.User.Uid}
	err = serv.UserRepo.First()
	if err != nil {
		return err
	}
	var user = &model.User{}
	if serv.UserRepo.User.Realname != userEdit.Realname {
		user.SetRealname(userEdit.Realname)
	}
	if serv.UserRepo.User.Qq != userEdit.Qq {
		user.SetQq(userEdit.Qq)
	}
	if serv.UserRepo.User.Mobile != userEdit.Mobile {
		user.SetMobile(userEdit.Mobile)
	}
	if serv.UserRepo.User.Email != userEdit.Email && !strutil.IsBlank(userEdit.Email) {
		//邮箱不为空 就检验邮箱验证码
		var emailCaptchaVerify = &requests.EmailCaptchaVerify{Email: userEdit.Email, Value: userEdit.Value}
		ok := verifyEmailCaptcha(emailCaptchaVerify)
		if !ok {
			err = exceptions.FailedVerify
			return
		}
		user.SetEmail(userEdit.Email)
	}
	if user != nil {
		user.Uid = global.User.Uid
		serv.UserRepo.User = user
		update, e := serv.UserRepo.Update()
		if e != nil {
			err = e
			return
		}
		if update < 1 {
			err = exceptions.ModifyError
			return
		}
	}
	return nil
}
func (serv *UserService) name() {
}
