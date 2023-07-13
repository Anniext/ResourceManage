package services

import (
	"ResourceManage/data"
	"ResourceManage/model"
	"ResourceManage/query"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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
	id := c.Query("id")
	var file *model.AvtFile
	var err string
	if id == "" {
		file, err = data.GetFile(name)

	} else {
		id, _ := strconv.ParseInt(id, 10, 64)
		query.AvtFile.Where(query.AvtFile.ID.Eq(id)).Scan(&file)
	}
	if err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		// 错误信息500,把error发送
		return
	}
	c.JSON(http.StatusOK, file)
}

func UnitGetGroup(c *gin.Context) {
	name := c.Query("name")
	id := c.Query("id")
	var unit *model.AvtUnit
	var err string
	if id == "" {
		unit, err = data.GetUnit(name)
	} else {
		id, _ := strconv.ParseInt(id, 10, 64)
		query.AvtUnit.Where(query.AvtUnit.ID.Eq(id)).Scan(&unit)
	}
	if err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		// 错误信息500,把error发送
		return
	}
	c.JSON(http.StatusOK, unit)
}

func UserGetGroup(c *gin.Context) {
	name := c.Query("name")
	unit, err := data.GetUser(name)
	if err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		// 错误信息500,把error发送
		return
	}
	c.JSON(http.StatusOK, unit)
}
