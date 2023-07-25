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

func (r *RouterGroup) Create(c *gin.Context) {
	groupName := GetGroupName(r)
	switch groupName {
	case RESOURCE:
		FileCreateGroup(c)
	case UNIT:
		UnitCreateGroup(c)
	case BACKEND:
		UserCreateGroup(c)
    case RELA:
        RelaCreateGroup(c)
	default:
		log.Println("Error group name", groupName)
	}
}

// FileCreateGroup 创建资源
func FileCreateGroup(c *gin.Context) {
	var file model.AvtFile
	if err := c.ShouldBind(&file); err != nil {
		c.JSON(http.StatusOK, api.JsonError(api.ErrCacheDate).JsonWithData(err))
		return
	}
	if ok, err := data.CreateFile(&file); !ok {
		c.JSON(http.StatusOK, api.JsonError(api.ErrCacheDate).JsonWithData(err))
		return
	}
	c.JSON(http.StatusOK, api.JsonSuccess())
}

func UnitCreateGroup(c *gin.Context) {
	var unit model.AvtUnit
	if err := c.ShouldBind(&unit); err != nil {
		c.JSON(http.StatusOK, api.JsonError(api.ErrCacheDate).JsonWithData(err))
		return
	}
	if ok, err := data.CreateUnit(&unit); !ok {
		c.JSON(http.StatusOK, api.JsonError(api.ErrCacheDate).JsonWithData(err))
		return
	}
	c.JSON(http.StatusOK, api.JsonSuccess())
}

func RelaCreateGroup(c *gin.Context) {
    var rela model.RelaUnitFile
	if err := c.ShouldBind(&rela); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if ok, err := data.CreateRela(&rela); !ok {
		c.JSON(http.StatusOK, api.JsonError(api.ErrCacheDate).JsonWithData(err))
		return
	}
	c.JSON(http.StatusOK, api.JsonSuccess())
}

func UserCreateGroup(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user model.SysBackendUser
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

