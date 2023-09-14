package service

import (
	"context"
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
