package main

import (
	"ResourceManage/config"
	"ResourceManage/data"
	services "ResourceManage/gin"
	"time"
)

func main() {
	//generate.GenGenerate()
	// 初始化配置文件
	config.Configinit()
	// 初始化服务器
	services.Serviceinit()
	time.Sleep(200 * time.Millisecond)
	data.SystemDataInit()
	//等待程序退出
	select {}
}
