package data

import "sync"

type BackendUserMap struct {
	data map[string]*SysBackendUser
	lock sync.RWMutex
}

func NewBackendUserMap() *BackendUserMap {
	return &BackendUserMap{
		data: make(map[string]*SysBackendUser),
	}
}
