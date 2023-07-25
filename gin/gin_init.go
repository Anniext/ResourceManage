package services

import (
	"ResourceManage/config"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

const (
    RESOURCE = "resource"
    UNIT = "unit"
    BACKEND = "user"
    RELA = "rela"
)

type HttpServerManager interface {
	SetHttp()
	Upload(c *gin.Context)
	Download(c *gin.Context)
}

type RouterManager interface {
	SetRouter()
	List(c *gin.Context)
	Get(c *gin.Context)
	Delete(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
}

type CacheInterface interface {
	SetCache()
	LoadFileData(c *gin.Context)
	LoadUnitData(c *gin.Context)
	LoadBackendUserData(c *gin.Context)
	LoadRelaUnitFileData(c *gin.Context)
}

type TokenInterface interface {
	SetToken()
	Login(c *gin.Context)
	Authentication(c *gin.Context)
}

var (
	router       *gin.Engine
	FileManager  RouterManager
	UnitManager  RouterManager
	UserManager  RouterManager
	RelaManager  RouterManager
	CacheManager CacheInterface
	HttpManager  HttpServerManager
	TokenManager TokenInterface
)

func RouterInit() {
	router = gin.Default()
	router.Use(DevCors()) //使用自定义的跨域中间件
	api := &RouterGroup{router.Group("/api")}
	FileManager = &RouterGroup{api.Group("/resource")}
	HttpManager = &RouterGroup{api.Group("/resource")}
	CacheManager = &RouterGroup{api.Group("/cache")}
	UnitManager = &RouterGroup{api.Group("/unit")}
	UserManager = &RouterGroup{api.Group("/user")}
	TokenManager = &RouterGroup{api.Group("/token")}
	RelaManager = &RouterGroup{api.Group("/rela")}
}

func Serviceinit() {
	}

func StartMainServer() {
	err := http.ListenAndServe(config.Configs.Dev.Router.Host, router)
	if err != nil {
		log.Println("Failed to start server: ", err)
	} else {
		log.Println("Server started on port: ", config.Configs.Dev.Router.Host)
		return
	}
}

func StartUploadServer() {
	err := http.ListenAndServe(config.Configs.Dev.Target.Upload, router)
	if err != nil {
		log.Println("Failed to start server: ", err)
	} else {
		log.Println("Server started on port: ", config.Configs.Dev.Target.Upload)
		return
	}
}

func StartDownloadServer() {
	err := http.ListenAndServe(config.Configs.Dev.Target.Download, router)
	if err != nil {
		log.Println("Failed to start server: ", err)
	} else {
		log.Println("Server started on port: ", config.Configs.Dev.Target.Download)
		return
	}
}

func GroupInit() {
	FileManager.SetRouter()
	UnitManager.SetRouter()
	UserManager.SetRouter()
	CacheManager.SetCache()
	HttpManager.SetHttp()
	TokenManager.SetToken()
	RelaManager.SetRouter()
}

func DevReturnDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Configs.Dev.DSN), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database:", err)
	}
	return db
}
