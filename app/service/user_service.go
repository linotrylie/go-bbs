package service

import (
	"context"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/random"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/gin-gonic/gin"
	"github.com/songzhibin97/gkit/net/ip"
	"go-bbs/app/exceptions"
	"go-bbs/app/http/model"
	"go-bbs/app/http/model/requests"
	"go-bbs/app/respository"
	"go-bbs/global"
	"go-bbs/utils"
	"time"
)

type userService struct {
}

var UserService = newUserService()

func newUserService() *userService {
	return new(userService)
}

// IsHasUserByUsername 是否有指定用户名的用户
// @return 存在 true  不存在 false
func (serv *userService) IsHasUserByUsername(username string, user *model.User) bool {
	where := make(map[string]interface{})
	where["username"] = username
	e := respository.UserRepository.GetDataByWhereMap(where)
	if e != nil {
		return false
	}
	if respository.UserRepository.User == nil {
		return false
	}
	user = respository.UserRepository.User
	return true
}

func (serv *userService) Login(userLogin requests.UserLogin, ctx *gin.Context) (userReturn *model.User, jwtCustomClaims *JwtCustomClaims, token string, e error) {
	//先检查是否存在相同用户名的用户
	user := model.User{Username: userLogin.Username}
	hasUser := serv.IsHasUserByUsername(userLogin.Username, &user)
	if !hasUser {
		return nil, nil, "", exceptions.UserNotFound
	}

	ok := serv.VerifyPassword(&user, userLogin.Password)
	if !ok {
		return nil, nil, "", exceptions.FailedVerify
	}
	//用户通过验证后，对用户进行后续操作，如增加经验积分或者记录登录ip等等
	go serv.LoginAfter(user, ctx)
	jwtCustomClaims, token, err := serv.ReturnUserInfo(*respository.UserRepository.User)
	if err != nil {
		return nil, nil, "", err
	}
	return &user, jwtCustomClaims, token, nil
}

func (serv *userService) GeneratePassword(user *model.User, password string) {
	var salt string
	if strutil.IsBlank(user.Salt) {
		salt = random.RandString(16)
	} else {
		salt = user.Salt
	}
	str := password + salt
	user.SetPassword(cryptor.Md5String(str)).SetSalt(salt)
}

func (serv *userService) VerifyPassword(user *model.User, password string) bool {
	str := password + user.Salt
	if cryptor.Md5String(str) == user.Password {
		return true
	}
	return false
}

func (serv *userService) LoginAfter(user model.User, ctx *gin.Context) {
	user.SetLogins(1).
		SetLoginDate(time.Now().Unix()).
		SetLoginIP(utils.Ip2long(ctx.ClientIP()))
	_, err := respository.UserRepository.Update(user)
	if err != nil {
		return
	}
}

func (serv *userService) Logout() {
	global.REDIS.Del(context.Background(), global.User.Username)
}

func (serv *userService) Detail(uid int) (*model.User, error) {
	user := model.User{Uid: uid}
	err := respository.UserRepository.FindByLocation(user)
	if err != nil {
		return nil, err
	}
	return respository.UserRepository.User, nil
}

func (serv *userService) ChangesPassword(userChangePassword *requests.UserChangePassword) (err error) {
	user := model.User{Uid: global.User.Uid}
	err = respository.UserRepository.FindByLocation(user)
	if err != nil {
		return
	}
	ok := serv.VerifyPassword(respository.UserRepository.User, userChangePassword.OldPassword)
	if !ok {
		err = exceptions.FailedVerify
		return
	}
	serv.GeneratePassword(respository.UserRepository.User, userChangePassword.NewPassword)
	update, err := respository.UserRepository.Update(user)
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

func (serv *userService) Edit(userEdit *requests.UserEdit) (err error) {
	user := model.User{Uid: global.User.Uid}
	err = respository.UserRepository.FindByLocation(user)
	if err != nil {
		return err
	}
	if respository.UserRepository.User.Realname != userEdit.Realname {
		user.SetRealname(userEdit.Realname)
	}
	if respository.UserRepository.User.Qq != userEdit.Qq {
		user.SetQq(userEdit.Qq)
	}
	if respository.UserRepository.User.Mobile != userEdit.Mobile {
		user.SetMobile(userEdit.Mobile)
	}
	if respository.UserRepository.User.Email != userEdit.Email && !strutil.IsBlank(userEdit.Email) {
		//邮箱不为空 就检验邮箱验证码
		var emailCaptchaVerify = &requests.EmailCaptchaVerify{Email: userEdit.Email, Value: userEdit.Value}
		ok := verifyEmailCaptcha(emailCaptchaVerify)
		if !ok {
			err = exceptions.FailedVerify
			return
		}
		user.SetEmail(userEdit.Email)
	}
	if &user != nil {
		var update, e = respository.UserRepository.Update(user)
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

func (serv *userService) Register(userRegister *requests.UserRegister, ctx *gin.Context) (*model.User, *JwtCustomClaims, string, error) {
	//先检查是否存在相同用户名的用户
	u := model.User{}
	hasUser := serv.IsHasUserByUsername(userRegister.Username, &u)
	if hasUser {
		return nil, nil, "", exceptions.DuplicateUser
	}
	//校验验证码
	//ok := ServiceGroupApp.CaptchaService.VerifyCaptcha(ctx, &userRegister.Captcha, &userRegister.Email)
	//if !ok {
	//	return nil, nil, "", exceptions.FailedVerify
	//}
	createIp, err := ip.StringToLong(ctx.ClientIP())
	var user = model.User{
		Username:   userRegister.Username,
		Email:      userRegister.Email.Email,
		CreateDate: time.Now().Unix(),
		CreateIp:   uint32(createIp),
		Gid:        101,
		Logins:     1,
		LoginDate:  time.Now().Unix(),
		LoginIp:    uint32(createIp),
		Signature:  "他什么也没留下~",
	}
	serv.GeneratePassword(&user, userRegister.Password)
	insert, err := respository.UserRepository.Insert(user)
	if err != nil {
		return nil, nil, "", err
	}
	if insert < 1 {
		return nil, nil, "", exceptions.CreateError
	}
	jwtCustomClaims, token, err := serv.ReturnUserInfo(*respository.UserRepository.User)
	if err != nil {
		return nil, nil, "", err
	}
	return &user, jwtCustomClaims, token, nil
}

func (serv *userService) ReturnUserInfo(user model.User) (jwtCustomClaims *JwtCustomClaims, token string, e error) {
	ServiceGroupApp.JwtService.SigningKey = []byte(global.CONFIG.JWT.SigningKey)
	claims := ServiceGroupApp.JwtService.CreateClaims(&user)
	token, e = ServiceGroupApp.JwtService.CreateToken(claims)
	if e != nil {
		return nil, "", e
	}
	global.User = respository.UserRepository.User
	jwtCustomClaims = &claims
	//将用户登录信息记录在redis中
	global.REDIS.Set(
		context.Background(), respository.UserRepository.User.Username,
		"login",
		time.Duration(
			utils.DatetimeToUnix(claims.ExpiresAt.Format(time.DateTime))-time.Now().Unix(),
		)*time.Second,
	)
	return
}

func (serv *userService) name() {
}
