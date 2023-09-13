package service

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/random"
	"go-bbs/app/exceptions"
	"go-bbs/app/http/model"
	"go-bbs/app/respository"
)

type UserService struct {
	userRepo respository.UserRepository
}

func (serv *UserService) Login(user *model.User) (e error) {
	//先检查是否存在相同用户名的用户
	serv.userRepo.User = &model.User{}
	where := make(map[string]interface{})
	where["username"] = user.Username
	e = serv.userRepo.FindUserByMap(where)
	if e != nil {
		return e
	}
	if serv.userRepo.User == nil {
		return exceptions.UserNotFound
	}
	ok := serv.VerifyPassword(serv.userRepo.User, user.Password)

	if !ok {
		return exceptions.FailedVerify
	}
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
	fmt.Println(user.Password)
	fmt.Println(password, str)
	fmt.Println(cryptor.Md5Byte([]byte(str)))
	data := []byte(str) //切片
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	fmt.Println(md5str)
	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	fmt.Printf("%s\n", hex.EncodeToString(cipherStr)) // 输出加密结果
	if cryptor.Md5String(str) == user.Password {
		return true
	}
	return false
}
