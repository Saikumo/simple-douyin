package util

import (
	"github.com/dgrijalva/jwt-go/v4"
	"saikumo.org/simple-douyin/src/entity"
	"time"
)

const issuer string = "Saikumo"

type Claims struct {
	UserId   int64
	Username string
	jwt.StandardClaims
}

func GenerateToken(user *entity.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(7 * 24 * time.Hour)
	claims := Claims{
		UserId:   user.Id,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(expireTime),
			Issuer:    issuer,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(issuer))
	return token, err
}
