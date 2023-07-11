package data

import (
	"ResourceManage/config"
	"log"
	"net/http"
)

var (
	CacheFile         *FileMap
	CacheUnit         *UnitMap
	CacheBackendUser  *BackendUserMap
	CacheRelaUnitFile *RelaUnitFileMap
)

func SystemDataInit() {
	// 建立缓存数据
	CacheFile = NewFileMap()
	CacheUnit = NewUnitMap()
	CacheBackendUser = NewBackendUserMap()
	CacheRelaUnitFile = NewRelaUnitFileMap()

	// 加载数据库数据
	err := GetFileData()
	if err != nil {
		return
	}
	err = GetUnitData()
	if err != nil {
		return
	}
	err = GetBackendUserData()
	if err != nil {
		return
	}
	err = GetRelaUnitFileData()
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

func GetBackendUserData() (err error) {
	_, err = http.Get("http://127.0.0.1" + config.Configs.AppPort + "/api/cache/LoadBackendUserData")
	if err != nil {
		log.Println("file_data请求失败:", err)
	}
	return
}
func GetRelaUnitFileData() (err error) {
	_, err = http.Get("http://127.0.0.1" + config.Configs.AppPort + "/api/cache/LoadRelaUnitFileData")
	if err != nil {
		log.Println("file_data请求失败:", err)
	}
	return
}
