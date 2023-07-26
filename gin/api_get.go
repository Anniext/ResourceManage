package services

import (
	"ResourceManage/api"
	"ResourceManage/data"
	"ResourceManage/model"
	"ResourceManage/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (r *RouterGroup) Get(c *gin.Context) {
	groupName := GetGroupName(r)
	switch groupName {
	case RESOURCE:
		FileGetGroup(c)
	case UNIT:
		UnitGetGroup(c)
	case BACKEND:
		UserGetGroup(c)
	case RELA:
		RelaGetGroup(c)
	default:
		log.Println("Error group name", groupName)
	}
}

func FileGetGroup(c *gin.Context) {
	name := c.Query("name")
	id := c.Query("id")
	queryType := c.Query("type")

	var file *model.AvtFile
	var err string

	// 猜猜为啥多写一道，嘿嘿，就是玩儿
	if query := utils.LimitParameter(name, id); query != "" {
		if queryType == "name" {
			id := data.CacheFile.GetID(query)
			file, err = data.GetFile(id)
		} else if queryType == "id" {
			id, _ := strconv.ParseInt(query, 10, 64)
			file, err = data.GetFile(id)
		} else {
			c.JSON(http.StatusOK, api.JsonError(api.ErrCacheDate).JsonWithData("参数错误"))
			return
		}
	}
	if err != "" {
		c.JSON(http.StatusOK, api.JsonError(api.ErrCacheDate).JsonWithData(err))
		return
	}
	c.JSON(http.StatusOK, api.JsonData(file))
}

func UnitGetGroup(c *gin.Context) {
	name := c.Query("name")
	id := c.Query("id")
	queryType := c.Query("type")

	var unit *model.AvtUnit
	var err string

	if query := utils.LimitParameter(name, id); query != "" {
		if queryType == "name" {
			unit, err = data.GetUnit(data.CacheUnit.GetID(query))
		} else if queryType == "id" {
			id, _ := strconv.ParseInt(query, 10, 64)
			unit, err = data.GetUnit(id)
		} else {
			c.JSON(http.StatusOK, api.JsonError(api.ErrCacheDate).JsonWithData("参数错误"))
			return
		}
	}
	if err != "" {
		c.JSON(http.StatusOK, api.JsonError(api.ErrCacheDate).JsonWithData(err))
		return
	}
	c.JSON(http.StatusOK, api.JsonData(unit))
}

func UserGetGroup(c *gin.Context) {
	name := c.Query("name")
	var user model.SysBackendUser
	if ok, err := data.GetUser(&user, name); !ok {
		c.JSON(http.StatusOK, api.JsonError(api.ErrCacheDate).JsonWithData(err))
		return
	}
	c.JSON(http.StatusOK, api.JsonData(user))
}

func RelaGetGroup(c *gin.Context) {
	id := c.Query("id")
	target := c.Query("target")
	idt, _ := strconv.ParseInt(id, 10, 64)
	re, err := data.GetRelaUnitFile(idt, target)
	if err != "" {
		c.JSON(http.StatusOK, api.JsonError(api.ErrCacheDate).JsonWithData(err))
		return
	}
	c.JSON(http.StatusOK, api.JsonData(re))
}
