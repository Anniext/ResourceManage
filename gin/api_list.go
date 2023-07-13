package services

import (
	"ResourceManage/data"
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
	arg := data.GetHeadBody{
		Page:   c.Param("page"),
		Limit:  c.Query("limit"),
		Offset: c.Query("offset"),
		Delete: c.Query("delete"),
	}
	list, count, err := data.GetFileList(&arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 错误信息500,把error发送
		return
	}
	c.JSON(http.StatusOK, gin.H{"filelist": list, "count": count})
}

func UnitListGroup(c *gin.Context) {
	arg := data.GetHeadBody{
		Page:   c.Param("page"),
		Limit:  c.Query("limit"),
		Offset: c.Query("offset"),
	}
	prmiss := c.MustGet("prmiss").(map[string]interface{})
	units, count, err := data.GetUnitList(&arg, prmiss)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 错误信息500,把error发送
		return
	}
	c.JSON(http.StatusOK, gin.H{"unitlist": units, "count": count})
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
