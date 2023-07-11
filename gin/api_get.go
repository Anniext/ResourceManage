package services

import (
	"ResourceManage/data"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func (r *RouterGroup) Get(c *gin.Context) {
	groupName := GetGroupName(r)
	switch groupName {
	case "resource":
		FileGetGroup(c)
	case "unit":
		UnitGetGroup(c)
	case "user":
		UserGetGroup(c)
	default:
		log.Println("Error group name", groupName)
	}
}

func FileGetGroup(c *gin.Context) {
	name := c.Query("name")
	file, err := data.GetFile(name)
	if err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		// 错误信息500,把error发送
		return
	}
	c.JSON(http.StatusOK, file)
}

func UnitGetGroup(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//获取路由参数
	id := c.Param("id")

	unit, err := data.GetUnit(id, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 错误信息500,把error发送
		return
	}

	c.JSON(http.StatusOK, unit)
}

func UserGetGroup(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//获取路由参数
	id := c.Param("id")
	unit, err := data.GetUser(id, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 错误信息500,把error发送
		return
	}
	c.JSON(http.StatusOK, unit)
}
