package services

import (
	"ResourceManage/data"
	"ResourceManage/token"
	"ResourceManage/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (r *RouterGroup) Login(c *gin.Context) {
	var user utils.UserInfo
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 错误信息400,把error发送
		return
	}
	if !CheckUser(user) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		// 错误信息401,把error发送
		return
	}

	Token, err := token.GetToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 错误信息500,把error发送
		return
	}
	// 发送状态码200
	c.JSON(http.StatusOK, gin.H{"msg": "Successful!", "token": Token})
}

// CheckUser 密码鉴权
func CheckUser(user utils.UserInfo) bool {
	logUserPwd := user.Password
	logUserPwdMd5 := utils.StringMd5(logUserPwd)
	userPwd := data.CacheBackendUser.Get(user.Username).UserPwd
	log.Println(userPwd)
	if logUserPwdMd5 == userPwd {
		return true
	} else {
		return false
	}
}
