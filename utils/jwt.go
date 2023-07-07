package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtUtil struct{}

var VerifyKey = []byte("Avtronsys")

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Prmisss  Permissions
}

type Permissions struct {
	Level   int `json:"level"`
	Expires int `json:"expires"`
}

type CustomClaims struct {
	*jwt.StandardClaims
	*UserInfo
}

func GeneratorJwt(userInfo *UserInfo) (string, error) {
	expire := userInfo.Prmisss.Expires
	claims := &CustomClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add((time.Hour * 24) * time.Duration(expire)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userInfo,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(VerifyKey)
}

func ParseCallBackJwt(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return VerifyKey, nil
}

func ParseJwt(tokenStr string) interface{} {
	token, err := jwt.Parse(tokenStr, ParseCallBackJwt)
	if token.Valid {
		return token.Claims
	}
	ve, ok := err.(*jwt.ValidationError)
	if ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return fmt.Sprintf("令牌解析失败: %v", ve)
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return fmt.Sprintf("令牌已经过期: %v", ve)
		} else {
			return fmt.Sprintf("无效令牌: %v", ve)
		}
	}
	return nil
}
