package services

import (
	"ResourceManage/api"
	"ResourceManage/data"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (r *RouterGroup) Delete(c *gin.Context) {
	groupName := GetGroupName(r)
	switch groupName {
	case RESOURCE:
		FileDeleteGroup(c)
	case UNIT:
		UnitDeleteGroup(c)
	case BACKEND:
		UserDeleteGroup(c)
    case RELA:
        RelaDeleteGroup(c)
	default:
		log.Println("Error group name", groupName)
	}
}

func FileDeleteGroup(c *gin.Context) {
	name := c.Query("name")
	if ok, err := data.DeleteFile(name); !ok {
		c.JSON(http.StatusOK, api.JsonError(api.ErrCacheDate).JsonWithData(err))
		return
	}
	c.JSON(http.StatusOK, api.JsonSuccess())
}

func UnitDeleteGroup(c *gin.Context) {
	name := c.Query("name")
    var prmiss interface {}
	if prmiss = c.MustGet("prmiss"); prmiss == nil {
		c.JSON(http.StatusOK, api.JsonError(api.JwtValidationErr).JsonWithData("令牌到期"))
		return
    }

	if ok, err := data.DeleteUnit(data.NamePrmiss{
		Name:   name,
		Primss: prmiss.(map[string]interface{}),
	}); !ok {
		c.JSON(http.StatusOK, api.JsonError(api.JwtValidationErr).JsonWithData(err))
		return
	}

	c.JSON(http.StatusOK, api.JsonSuccess())
}

func RelaDeleteGroup(c *gin.Context) {
	name := c.Query("unit")
	file := c.Query("file")

	if ok, err := data.DeleteRela(name, file); !ok {
		c.JSON(http.StatusOK, api.JsonError(api.ErrCacheDate).JsonWithData(err))
		return
	}
	c.JSON(http.StatusOK, api.JsonSuccess())
}

func UserDeleteGroup(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	if err := data.DeleteUser(id, db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "User deleted successfully"})
}
