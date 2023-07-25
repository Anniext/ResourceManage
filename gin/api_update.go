package services

import (
	"ResourceManage/api"
	"ResourceManage/data"
	"ResourceManage/model"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (r *RouterGroup) Update(c *gin.Context) {
	groupName := GetGroupName(r)
	switch groupName {
	case RESOURCE:
		FileUpdateGroup(c)
	case UNIT:
		UnitUpdateGroup(c)
	case BACKEND:
		UserUpdateGroup(c)
	case RELA:
		RelaUpdateGroup(c)
	default:
		log.Println("Error group name", groupName)
	}
}

func FileUpdateGroup(c *gin.Context) {
	name := c.Query("name")
	var file model.AvtFile
	if err := c.ShouldBind(&file); err != nil {
		c.JSON(http.StatusOK, api.JsonError(api.ErrServer).JsonWithData(err))
		return
	}
	if ok, err := data.UpdateFile(name, &file); !ok {
		c.JSON(http.StatusOK, api.JsonError(api.ErrCacheDate).JsonWithData(err))
		return
	}
	c.JSON(http.StatusOK, api.JsonSuccess())
}

func UnitUpdateGroup(c *gin.Context) {
	name := c.Query("name")
	var unit model.AvtUnit
	if err := c.ShouldBind(&unit); err != nil {
		c.JSON(http.StatusOK, api.JsonError(api.ErrServer).JsonWithData(err))
		return
	}
	if ok, err := data.UpdateUnit(name, &unit); !ok {
		c.JSON(http.StatusOK, api.JsonError(api.ErrCacheDate).JsonWithData(err))
		return
	}
	c.JSON(http.StatusOK, api.JsonSuccess())
}

func RelaUpdateGroup(c *gin.Context) {
	name := c.Query("name")
	var rela model.RelaUnitFile
	if err := c.ShouldBind(&rela); err != nil {
		c.JSON(http.StatusOK, api.JsonError(api.ErrServer).JsonWithData(err))
		return
	}
	if ok, err := data.UpdateRela(name, &rela); !ok {
		c.JSON(http.StatusOK, api.JsonError(api.ErrCacheDate).JsonWithData(err))
		return
	}
	c.JSON(http.StatusOK, api.JsonSuccess())
}

func UserUpdateGroup(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var user model.SysBackendUser
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, api.JsonError(api.ErrServer).JsonWithData(err))
		return
	}
	if err := data.UpdateUser(id, &user, db); err != "" {
		c.JSON(http.StatusOK, api.JsonError(api.ErrCacheDate).JsonWithData(err))
		return
	}
	c.JSON(http.StatusOK, api.JsonSuccess())
}
