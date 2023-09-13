package service

import (
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/random"
	"github.com/gin-gonic/gin"
	"go-bbs/app/exceptions"
	"go-bbs/app/http/model"
	"go-bbs/app/respository"
	"go-bbs/global"
	"go-bbs/utils"
	"time"
)

type UserService struct {
	userRepo respository.UserRepository
}

func (serv *UserService) Login(user *model.User, ctx *gin.Context) (userReturn *model.User, token string, e error) {
	//先检查是否存在相同用户名的用户
	serv.userRepo.User = &model.User{}
	where := make(map[string]interface{})
	where["username"] = user.Username
	e = serv.userRepo.FindUserByMap(where)
	if e != nil {
		return nil, "", e
	}
	if serv.userRepo.User == nil {
		return nil, "", exceptions.UserNotFound
	}
	ok := serv.VerifyPassword(serv.userRepo.User, user.Password)
	if !ok {
		return nil, "", exceptions.FailedVerify
	}
	ServiceGroupApp.JwtService.SigningKey = []byte(global.CONFIG.JWT.SigningKey)
	claims := ServiceGroupApp.JwtService.CreateClaims(serv.userRepo.User)
	token, e = ServiceGroupApp.JwtService.CreateToken(claims)
	if e != nil {
		return nil, "", e
	}
	userReturn = serv.userRepo.User
	global.User = serv.userRepo.User
	//用户通过验证后，对用户进行后续操作，如增加经验积分或者记录登录ip等等
	go serv.LoginAfter(ctx)
	return
}

func (serv *UserService) GeneratePassword(user *model.User, password string) {
	salt := random.RandString(16)
	str := password + salt
	user.Password = cryptor.Md5String(str)
	user.Salt = salt
}

func (serv *UserService) name() {
}

func (serv *UserService) VerifyPassword(user *model.User, password string) bool {
	str := password + user.Salt
	if cryptor.Md5String(str) == user.Password {
		return true
	}
	return false
}

func (serv *UserService) LoginAfter(ctx *gin.Context) {
	serv.userRepo.User.SetLogins(1).
		SetLoginDate(int(time.Now().Unix())).
		SetLoginIP(int(utils.Ip2long(ctx.ClientIP())))
	_, err := serv.userRepo.Update()
	if err != nil {
		return
	}
}
