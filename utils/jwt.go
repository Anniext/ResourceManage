package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

type JwtUtil struct{}

var VerifyKey = []byte("Avtronsys")

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CustomClaims struct {
	*jwt.StandardClaims
	*UserInfo
}

func GeneratorJwt(userInfo *UserInfo) (string, error) {
	claims := &CustomClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userInfo,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(VerifyKey)
}

func ParseJwt(tokenStr string) *jwt.Token {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return VerifyKey, nil
	})

	if token.Valid {
		log.Println("You look nice today")
		return token
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			log.Printf("%v", ve)
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			log.Printf("%v", ve)
			panic("无效令牌, 已经过期")
		} else {
			panic("无效令牌")
		}
	} else {
		panic("无效令牌")
	}
	return token
}
