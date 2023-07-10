package services

import (
	"ResourceManage/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *RouterGroup) LoadFileData(c *gin.Context) {
	err := data.LoadFileData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, err)
}

func (r *RouterGroup) LoadUnitData(c *gin.Context) {
	err := data.LoadUnitData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, err)
}

func (r *RouterGroup) LoadBackendUserData(c *gin.Context) {
	err := data.LoadBackendUserData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, err)
}
