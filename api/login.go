package services

import (
	"ResourceManage/data"
	"ResourceManage/token"
	"ResourceManage/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *RouterGroup) Login(c *gin.Context) {
	var user utils.UserInfo
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 错误信息400,把error发送
		return
	}
	if !CheckUser(&user) {
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

func (r *RouterGroup) Authentication(c *gin.Context) {
	tk := c.Request.Header.Get("Authorization")
	if tk == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token为空"})
		return
	}
	result := utils.ParseJwt(tk)
	if result == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token无效"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": result})
}

// CheckUser 密码鉴权
func CheckUser(user *utils.UserInfo) bool {
	logUserPwd := user.Password
	logUserPwdMd5 := utils.StringMd5(logUserPwd)
	sysUser := data.CacheBackendUser.Get(user.Username)
	userPwd := sysUser.UserPwd
	if logUserPwdMd5 == userPwd {
		user.Prmisss.Expires = sysUser.Expires
		user.Prmisss.Level = sysUser.Level
		return true
	} else {
		return false
	}
}
