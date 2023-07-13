package services

import (
	"ResourceManage/data"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func (r *RouterGroup) Delete(c *gin.Context) {
	groupName := GetGroupName(r)
	switch groupName {
	case "resource":
		FileDeleteGroup(c)
	case "unit":
		UnitDeleteGroup(c)
	case "user":
		UserDeleteGroup(c)
	default:
		log.Println("Error group name", groupName)
	}
}

func FileDeleteGroup(c *gin.Context) {
	name := c.Query("name")
	if err := data.DeleteFile(name); err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "File deleted successfully"})
}

func UnitDeleteGroup(c *gin.Context) {
	name := c.Query("name")
	prmiss := c.MustGet("prmiss")

	if err := data.DeleteUnit(data.NamePrmiss{
		Name:   name,
		Primss: prmiss.(map[string]interface{}),
	}); err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Unit deleted successfully"})
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
