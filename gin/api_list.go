package services

import (
	"ResourceManage/api"
	"ResourceManage/data"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func (r *RouterGroup) List(c *gin.Context) {
	groupName := GetGroupName(r)
	switch groupName {
	case RESOURCE:
		FileListGroup(c)
	case UNIT:
		UnitListGroup(c)
	case BACKEND:
		UserListGroup(c)
	case RELA:
		RelaListGroup(c)
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
	filelist := data.GetFileList(&arg).(data.FileList)
	if filelist.Error != nil {
		c.JSON(http.StatusOK, api.JsonError(api.ErrCacheDate).JsonWithData(filelist.Error))
		return
	}
	c.JSON(http.StatusOK, api.JsonData(filelist))
}

func UnitListGroup(c *gin.Context) {
	arg := data.GetHeadBody{
		Page:   c.Param("page"),
		Limit:  c.Query("limit"),
		Offset: c.Query("offset"),
	}
	prmiss := c.MustGet("prmiss").(map[string]interface{})
	unitlist := data.GetUnitList(&arg, prmiss).(data.UnitList)
	if unitlist.Error != nil {
		c.JSON(http.StatusOK, api.JsonError(api.ErrCacheDate).JsonWithData(unitlist.Error))
		return
	}
	c.JSON(http.StatusOK, api.JsonData(unitlist))
}

func RelaListGroup(c *gin.Context) {
	arg := data.GetHeadBody{
		Page:   c.Param("page"),
		Limit:  c.Query("limit"),
		Offset: c.Query("offset"),
	}
	id := c.Query("id")
	target := c.Query("target")
	idt, _ := strconv.ParseInt(id, 10, 64)

	list := data.GetRelaList(&arg, idt, target).(data.RelaList)
	if list.Error != nil {
		c.JSON(http.StatusOK, api.JsonError(api.ErrCacheDate).JsonWithData(list.Error))
		return
	}
	c.JSON(http.StatusOK, api.JsonData(list))
}

func UserListGroup(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	list, err := data.GetUserList(db)
	if err != nil {
		c.JSON(http.StatusOK, api.JsonError(api.ErrCacheDate).JsonWithData(err))
		return
	}
	c.JSON(http.StatusOK, api.JsonData(
		gin.H{
			"filelist": list,
		}),
	)
}
