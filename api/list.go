package services

import (
	"ResourceManage/data"
	"ResourceManage/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"regexp"
)

func (r *RouterGroup) List(c *gin.Context) {
	groupName := GetGroupName(r)
	switch groupName {
	case "resource":
		FileListGroup(c)
	case "unit":
		UnitListGroup(c)
	case "user":
		UserListGroup(c)
	default:
		log.Println("Error group name", groupName)
	}
}

func GetGroupName(r *RouterGroup) string {
	basePath := r.BasePath()
	re := regexp.MustCompile(`/api/(?P<match>[^/]+)`)
	match := re.FindStringSubmatch(basePath)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func FileListGroup(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	delete_id := c.Param("delete_id")
	list, err := data.GetFileList(db, delete_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 错误信息500,把error发送
		return
	}
	c.JSON(http.StatusOK, list)
}

func UnitListGroup(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	delete_id := c.Param("delete_id")
	level, errStr := utils.GetLevel(utils.GetJwtClaims(c)) //通过token获取level
	if errStr != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errStr})
		return
	}
	units, err := data.GetUnitList(db, delete_id, level)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 错误信息500,把error发送
		return
	}
	c.JSON(http.StatusOK, units)
}

func UserListGroup(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	units, err := data.GetUserList(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 错误信息500,把error发送
		return
	}
	c.JSON(http.StatusOK, units)
}
