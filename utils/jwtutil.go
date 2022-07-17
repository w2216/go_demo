package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type Users struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CustomClaims struct {
	Users
	jwt.StandardClaims
}

type WebUsers struct {
	Id     uint
	Email  string
	Params map[string]interface{}
}

type WebClaims struct {
	WebUsers
	jwt.StandardClaims
}

var MySecret = []byte("密钥")

// 创建 Token
func GenToken(user Users) (string, error) {
	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Minute * 5)), //5分钟后过期
			Issuer:    "xx",                                    //签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}

// 创建 WebToken
func WebToken(webUsers WebUsers) (string, error) {
	claim := WebClaims{
		webUsers,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Minute * 5)), //5分钟后过期
			Issuer:    "xx",                                    //签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}

// 解析 token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		fmt.Println(" token parse err:", err)
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// 刷新 Token
func RefreshToken(tokenStr string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = jwt.At(time.Now().Add(time.Minute * 10))
		return GenToken(claims.Users)
	}
	return "", errors.New("cloudn't handle this token")
}
