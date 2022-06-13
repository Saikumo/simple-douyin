package util

import (
	"github.com/dgrijalva/jwt-go/v4"
	"saikumo.org/simple-douyin/src/common"
	"saikumo.org/simple-douyin/src/entity"
	"time"
)

const ISSUER string = "Saikumo"

var secret []byte = []byte("Saikumo")

type Claims struct {
	UserId   uint
	Username string
	jwt.StandardClaims
}

func GenerateToken(user *entity.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(7 * 24 * time.Hour)
	claims := Claims{
		UserId:   user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(expireTime),
			Issuer:    ISSUER,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	return token, err
}

func CheckToken(token string) (*Claims, error) {
	//解析token
	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	//token无效
	if !parsedToken.Valid {
		return nil, common.TokenIsNotValid
	}
	//取出claim 校验过期时间
	claims := parsedToken.Claims.(*Claims)
	if time.Now().Unix() > claims.ExpiresAt.Unix() {
		return nil, common.TokenIsExpired
	}
	return claims, nil
}
