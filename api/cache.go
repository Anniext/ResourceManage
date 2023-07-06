package services

import (
	"ResourceManage/data"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (r *RouterGroup) LoadFileData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	data.LoadFileData(db)
	c.JSON(200, nil)
}

func (r *RouterGroup) LoadUnitData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	data.LoadUnitData(db)
	c.JSON(200, nil)
}

func (r *RouterGroup) LoadUserData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	data.LoadUserData(db)
	c.JSON(200, nil)
}
func (r *RouterGroup) LoadBackendUserData(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	data.LoadBackendUserData(db)
	c.JSON(200, nil)
}
