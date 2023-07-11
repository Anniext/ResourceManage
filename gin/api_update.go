package services

import (
	"ResourceManage/data"
	"ResourceManage/model"
	"ResourceManage/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func (r *RouterGroup) Update(c *gin.Context) {
	groupName := GetGroupName(r)
	switch groupName {
	case "resource":
		FileUpdateGroup(c)
	case "unit":
		UnitUpdateGroup(c)
	case "user":
		UserUpdateGroup(c)
	default:
		log.Println("Error group name", groupName)
	}
}

func FileUpdateGroup(c *gin.Context) {
	name := c.Query("name")
	var file model.AvtFile
	if err := c.ShouldBind(&file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 错误信息400,把error发送
		return
	}
	if err := data.UpdateFile(name, &file); err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		// 错误信息500,把error发送
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "File updated successfully"})
}

func UnitUpdateGroup(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var unit model.AvtUnit
	if err := c.ShouldBind(&unit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 错误信息400,把error发送
		return
	}
	level, errStr := utils.GetLevel(utils.GetJwtClaims(c)) //通过token获取level
	if errStr != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errStr})
		return
	}
	if unit.Level <= level {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Permission too low to create"})
		return
	}
	if err := data.UpdateUnit(id, &unit, db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 错误信息500,把error发送
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Unit updated successfully"})
}

func UserUpdateGroup(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var user model.SysBackendUser
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 错误信息400,把error发送
		return
	}
	if err := data.UpdateUser(id, &user, db); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 错误信息500,把error发送
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
