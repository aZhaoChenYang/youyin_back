package common

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

var Secret = []byte("secret")

const TokenExpireDuration = time.Hour * 24

// GenToken 生成Token
func GenToken(userId string) (string, error) {
	claims := MyClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "gin-blog",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Secret)
}

// CheckToken 验证Token
func CheckToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// ParseToken 解析token
//func ParseToken(tokenString string) (*Myclaims, error) {
//	token, err := jwt.ParseWithClaims(tokenString, &Myclaims{}, func(token *jwt.Token) (interface{}, error) {
//		return Secret, nil
//	})
//	if err != nil {
//		return nil, err
//	}
//	if claims, ok := token.Claims.(*Myclaims); ok && token.Valid {
//		return claims, nil
//	}
//	return nil, errors.New("invalid token")
//}
//// GenToken 生成Token
//func GenToken(username string, userId int) (string, error) {
//	c := Myclaims{
//		Username: username,
//		UserId:   userId,
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
//			Issuer:    "gin-demo",
//		},
//	}
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
//	return token.SignedString(Secret)
//}
