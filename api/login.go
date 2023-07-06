package services

import (
	"ResourceManage/token"
	"ResourceManage/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *RouterGroup) Login(c *gin.Context) {
	//db := c.MustGet("db").(*gorm.DB)
	var user utils.UserInfo
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 错误信息400,把error发送
		return
	}

	Token, err := token.GetToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 错误信息500,把error发送
		return
	}
	// 发送状态码200
	c.JSON(http.StatusOK, gin.H{"message": Token})

}
