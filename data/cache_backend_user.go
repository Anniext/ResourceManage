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
	err = query.SysBackendUser.Scan(&userDataList)
	if err != nil {
		log.Println("sys_backend_user表数据加载错误：", err)
		return err
	}
	count, _ := query.SysBackendUser.Count()
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

func (m *BackendUserMap) Update(user *model.SysBackendUser) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if _, ok := m.data[user.UserName]; ok {
		delete(m.data, user.UserName)
		m.data[user.UserName] = user
	}
	m.data[user.UserName] = user
	_, err := query.SysBackendUser.Where(query.SysBackendUser.ID.Eq(user.ID)).Updates(map[string]interface{}{
		"user_name": user.UserName,
		"status":    user.Status,
		"email":     user.Email,
		"level":     user.Level,
	})
	if err != nil {
		log.Println("avt_file表数据更新错误：", err)
		return
	}
}
