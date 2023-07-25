package main

import (
	"ResourceManage/config"
	"ResourceManage/data"
	services "ResourceManage/gin"
	"ResourceManage/query"
	"time"
)

const (
	APPMODE = "dev"
	NULL    = ""
)

func main() {
	//generate.GenGenerate()
	config.Configinit()
	if config.Configs.Dev.DSN != NULL {
		query.SetDefault(services.DevReturnDB()) //初始化数据服务
	}

	if config.Configs.Mode == APPMODE {
		services.RouterInit()
		services.GroupInit()              // 初始化路由
		go services.StartDownloadServer() // 启动下载服务
	}

	if config.Configs.Dev.Target.Download != NULL && config.Configs.Dev.Target.Upload != NULL {
		go services.StartMainServer()   // 启动主服务
		go services.StartUploadServer() // 启动上传服务
	}

	time.Sleep(200 * time.Millisecond)
	data.SystemDataInit()
	//等待程序退出
	select {}
}
