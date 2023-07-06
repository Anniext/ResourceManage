package token

import (
	"ResourceManage/utils"
	"log"
)

func GetToken(user utils.UserInfo) (string, error) {
	token, err := utils.GeneratorJwt(&user)
	if err != nil {
		log.Println(err)
		return "token申请失败", err
	}
	//re := utils.ParseJwt(token)
	return token, nil
}
