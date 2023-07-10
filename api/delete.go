package services

import (
	"ResourceManage/data"
	"ResourceManage/utils"
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
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	if err := data.DeleteFile(id, db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File deleted"})
}

func UnitDeleteGroup(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	level, errStr := utils.GetLevel(utils.GetJwtClaims(c)) //通过token获取level
	if errStr != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errStr})
		return
	}
	unit, _ := data.GetUnit(id, db)
	if unit.Level <= level {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Permission too low to delete"})
		return
	}
	if err := data.DeleteUnit(id, db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Unit deleted"})
}

func UserDeleteGroup(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	if err := data.DeleteUser(id, db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
