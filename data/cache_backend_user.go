package data

import (
	"ResourceManage/model"
	"ResourceManage/query"
	"log"
	"sync"
)

type BackendUserMap struct {
	data map[string]*model.SysBackendUser
	lock sync.RWMutex
}

func NewBackendUserMap() *BackendUserMap {
	return &BackendUserMap{
		data: make(map[string]*model.SysBackendUser),
	}
}

func LoadBackendUserData() (err error) {
	var userDataList []*model.SysBackendUser
	err = query.AvtFile.Scan(&userDataList)
	if err != nil {
		log.Println("sys_backend_user表数据加载错误：", err)
		return err
	}
	count, _ := query.AvtFile.Count()
	if count > 0 {
		log.Println("sys_backend_user表缓存数据加载成功!")
		for _, user := range userDataList {
			CacheBackendUser.Set(user)
		}
	} else {
		log.Println("sys_backend_user没有数据！")
		return err
	}
	return nil
}
func (m *BackendUserMap) Set(bu *model.SysBackendUser) {
	m.lock.Lock()
	m.data[bu.UserName] = bu
	m.lock.Unlock()
}

func (m *BackendUserMap) Get(name string) *model.SysBackendUser {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.data[name]
}
