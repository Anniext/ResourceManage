package data

import (
	"ResourceManage/config"
	"log"
	"net/http"
)

var (
	CacheFile *FileMap
	CacheUnit *UnitMap
	CacheUser *UserMap
)

func SystemDataInit() {
	// 建立缓存数据
	CacheFile = NewFileMap()
	CacheUnit = NewUnitMap()
	CacheUser = NewUserMap()

	// 加载数据库数据
	err := GetFileData()
	if err != nil {
		return
	}
	err = GetUnitData()
	if err != nil {
		return
	}
	err = GetUserData()
	if err != nil {
		return
	}
}

func GetFileData() (err error) {
	_, err = http.Get("http://127.0.0.1" + config.Configs.AppPort + "/api/cache/LoadFileData")
	if err != nil {
		log.Println("file_data请求失败:", err)
	}
	return
}

func GetUnitData() (err error) {
	_, err = http.Get("http://127.0.0.1" + config.Configs.AppPort + "/api/cache/LoadUnitData")
	if err != nil {
		log.Println("file_data请求失败:", err)
	}
	return
}

func GetUserData() (err error) {
	_, err = http.Get("http://127.0.0.1" + config.Configs.AppPort + "/api/cache/LoadUserData")
	if err != nil {
		log.Println("file_data请求失败:", err)
	}
	return
}
