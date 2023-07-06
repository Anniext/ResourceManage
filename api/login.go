package services

import (
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
	// 从缓存中获取数据判断是否存在数据

	Token, err := token.GetToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 错误信息500,把error发送
		return
	}
	// 发送状态码200
	c.JSON(http.StatusOK, gin.H{"message": Token})

}

//
//func CheckPassWd() {
//
//}
