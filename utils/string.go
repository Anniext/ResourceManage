package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

func StringSplit(s string, sep string) []string {
	return strings.Split(s, sep)
}

func StringMd5(s string) string {
	hash := md5.Sum([]byte(s))                // 使用md5.Sum对字符串进行哈希
	hashString := hex.EncodeToString(hash[:]) // 将哈希值转换为十六进制字符串
	return hashString
}

func GetLevel(i interface{}) (prmiss map[string]interface{}, err string) {
	err, ok := i.(string)
	if ok {
		return nil, err
	}
	prmiss = i.(jwt.MapClaims)["Prmiss"].(map[string]interface{})
	return
}
