package services

import (
	"github.com/gin-gonic/gin"
	"log"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func (r *RouterGroup) SetRouter() {
	groupName := GetGroupName(r)
	switch groupName {
	case "resource":
		SetResource(r)
	case "unit":
		SetUnit(r)
	case "user":
		SetUser(r)
	default:
		log.Println("Error group name", groupName)
	}
}

func SetResource(r *RouterGroup) {
	r.POST("/create", FileManager.Create)
	r.GET("/get/:id", FileManager.Get)
	r.GET("/list/:delete_id", FileManager.List)
	r.PUT("/update/:id", FileManager.Update)
	r.DELETE("/delete/:id", FileManager.Delete)

}

func SetUnit(r *RouterGroup) {
	r.POST("/create", UnitManager.Create)
	r.GET("/get/:id", UnitManager.Get)
	r.GET("/list/:delete_id", UnitManager.List)
	r.PUT("/update/:id", UnitManager.Update)
	r.DELETE("/delete/:id", UnitManager.Delete)
}

func SetUser(r *RouterGroup) {
	r.POST("/create", UserManager.Create)
	r.GET("/get/:id", UserManager.Get)
	r.GET("/list/:delete_id", UserManager.List)
	r.PUT("/update/:id", UserManager.Update)
	r.DELETE("/delete/:id", UserManager.Delete)
}

func (r *RouterGroup) SetCache() {
	r.GET("/LoadFileData", CacheManager.LoadFileData)
	r.GET("/LoadUnitData", CacheManager.LoadUnitData)
	r.GET("/LoadUserData", CacheManager.LoadUnitData)
	r.GET("/LoadBackendUserData", CacheManager.LoadBackendUserData)
}

func (r *RouterGroup) SetHttp() {
	r.POST("/upload", HttpManager.Upload)
	r.POST("/download", HttpManager.Download)
}

func (r *RouterGroup) SetToken() {
	r.POST("/login", TokenManager.Login)
}
