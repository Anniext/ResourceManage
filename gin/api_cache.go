package services

import (
	"ResourceManage/api"
	"ResourceManage/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *RouterGroup) LoadFileData(c *gin.Context) {
	err := data.LoadFileData()
	if err != nil {
        c.JSON(http.StatusOK,api.JsonError(api.ErrCache).JsonWithData(err))
	}
	c.JSON(http.StatusOK, api.JsonSuccess())
}

func (r *RouterGroup) LoadUnitData(c *gin.Context) {
	err := data.LoadUnitData()
	if err != nil {
        c.JSON(http.StatusOK,api.JsonError(api.ErrCache).JsonWithData(err))
	}
	c.JSON(http.StatusOK, api.JsonSuccess())
}

func (r *RouterGroup) LoadBackendUserData(c *gin.Context) {
	err := data.LoadBackendUserData()
	if err != nil {
        c.JSON(http.StatusOK,api.JsonError(api.ErrCache).JsonWithData(err))
	}
	c.JSON(http.StatusOK, api.JsonSuccess())
}

func (r *RouterGroup) LoadRelaUnitFileData(c *gin.Context) {
	err := data.LoadRelaUnitFileData()
	if err != nil {
        c.JSON(http.StatusOK,api.JsonError(api.ErrCache).JsonWithData(err))
	}
	c.JSON(http.StatusOK, api.JsonSuccess())
}
