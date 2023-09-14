package service

import (
	jwt "github.com/golang-jwt/jwt/v5"
	"go-bbs/app/exceptions"
	"go-bbs/app/http/model"
	"go-bbs/global"
	"go-bbs/utils"
	"time"
)

type JwtService struct {
	SigningKey []byte
}

type JwtCustomClaims struct {
	UID        int
	GID        int
	Email      string
	Username   string
	Mobile     string
	BufferTime int64
	jwt.RegisteredClaims
}

func (serv *JwtService) CreateClaims(user *model.User) JwtCustomClaims {
	bf, _ := utils.ParseDuration(global.CONFIG.JWT.BufferTime)
	ep, _ := utils.ParseDuration(global.CONFIG.JWT.ExpiresTime)
	jwtCustomClaims := JwtCustomClaims{
		UID:        user.Uid,
		GID:        user.Gid,
		Email:      user.Email,
		Username:   user.Username,
		Mobile:     user.Mobile,
		BufferTime: int64(bf / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    global.CONFIG.JWT.Issuer,
			Subject:   "Token",
			Audience:  jwt.ClaimStrings{"FREE-BNS"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)),
		},
	}
	return jwtCustomClaims
}

func (serv *JwtService) CreateToken(claims JwtCustomClaims) (token string, err error) {
	if serv.SigningKey == nil {
		return "", nil
	}
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(serv.SigningKey)
	return
}

func (serv *JwtService) CreateTokenByOldToken(oldToken string, claims JwtCustomClaims) (string, error) {
	v, err, _ := global.Sf.Do("JWT:"+oldToken, func() (interface{}, error) {
		return serv.CreateToken(claims)
	})
	return v.(string), err
}

// ParseToken 解析 token
func (serv *JwtService) ParseToken(tokenString string) (*JwtCustomClaims, error) {
	jwtCustomClaims := &JwtCustomClaims{}
	serv.SigningKey = []byte(global.CONFIG.JWT.SigningKey)
	token, err := jwt.ParseWithClaims(tokenString, jwtCustomClaims, func(token *jwt.Token) (i interface{}, e error) {
		return serv.SigningKey, nil
	})
	if err != nil && token.Valid {
		err = exceptions.TokenInvalid
	}
	return jwtCustomClaims, err
}

// IsTokenValid 判断token是否有效
func (serv *JwtService) IsTokenValid(tokenStr string) bool {
	_, err := serv.ParseToken(tokenStr)
	if err != nil {
		return false
	}
	return true
}
