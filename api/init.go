package services

import (
	"ResourceManage/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
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
	LoadUserData(c *gin.Context)
	LoadBackendUserData(c *gin.Context)
}

type TokenInterface interface {
	SetToken()
	Login(c *gin.Context)
	Authentication(c *gin.Context)
}

var (
	router       *gin.Engine
	api          *RouterGroup
	FileManager  RouterManager
	UnitManager  RouterManager
	UserManager  RouterManager
	CacheManager CacheInterface
	HttpManager  HttpServerManager
	TokenManager TokenInterface
)

func init() {
	RouterGroupInit()
}

func RouterGroupInit() {
	//gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	db := SqlserverInit()
	router.Use(cors.Default())
	api = &RouterGroup{router.Group("/api")}
	api.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	FileManager = &RouterGroup{api.Group("/resource")}
	HttpManager = &RouterGroup{api.Group("/resource")}
	CacheManager = &RouterGroup{api.Group("/cache")}
	UnitManager = &RouterGroup{api.Group("/unit")}
	UserManager = &RouterGroup{api.Group("/user")}
	TokenManager = &RouterGroup{api.Group("/token")}
}

func Serviceinit() {
	GroupInit()              // 初始化路由
	go StartMainServer()     // 启动主服务
	go StartUploadServer()   // 启动上传服务
	go StartDownloadServer() // 启动下载服务
}

func StartMainServer() {
	err := http.ListenAndServe(config.Configs.AppPort, router)
	if err != nil {
		log.Println("Failed to start server: ", err)
	} else {
		log.Println("Server started on port: ", config.Configs.AppPort)
		return
	}
}

func StartUploadServer() {
	err := http.ListenAndServe(config.Configs.UploadPort, router)
	if err != nil {
		log.Println("Failed to start server: ", err)
	} else {
		log.Println("Server started on port: ", config.Configs.UploadPort)
		return
	}
}

func StartDownloadServer() {
	err := http.ListenAndServe(config.Configs.DownloadPort, router)
	if err != nil {
		log.Println("Failed to start server: ", err)
	} else {
		log.Println("Server started on port: ", config.Configs.DownloadPort)
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
}

func SqlserverInit() (db *gorm.DB) {
	log.Println("mysql connection...", config.Configs.Dsn)
	db, err := gorm.Open(mysql.Open(config.Configs.Dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database:", err)
	} else {
		log.Println("mysql connection successful!")
	}
	if err != nil {
		panic("Failed to get underlying DB connection")
	}
	return db
}
