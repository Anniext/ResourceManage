package services

import (
	"ResourceManage/data"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func (r *RouterGroup) Create(c *gin.Context) {
	groupName := GetGroupName(r)
	switch groupName {
	case "resource":
		FileCreateGroup(c)
	case "unit":
		UnitCreateGroup(c)
	case "user":
		UserCreateGroup(c)
	default:
		log.Println("Error group name", groupName)
	}
}

func FileCreateGroup(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var file data.AvtFile
	// 请求响应绑定File结构
	if err := c.ShouldBind(&file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 错误信息400,把error发送
		return
	}
	//log.Println(r)
	if err := data.CreateFile(&file, db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 错误信息500,把error发送
		return
	}
	// 发送状态码200
	c.JSON(http.StatusOK, "File created successfully")
}

func UnitCreateGroup(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var unit data.AvtUnit
	// 请求响应绑定File结构
	if err := c.ShouldBind(&unit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 错误信息400,把error发送
		return
	}
	if err := data.CreateUnit(&unit, db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 错误信息500,把error发送
		return
	}
	// 发送状态码200
	c.JSON(http.StatusOK, "Unit created successfully")
}

func UserCreateGroup(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user data.AvtUser
	// 请求响应绑定File结构
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 错误信息400,把error发送
		return
	}
	if err := data.CreateUser(&user, db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 错误信息500,把error发送
		return
	}
	// 发送状态码200
	c.JSON(http.StatusOK, "User created successfully")
}
