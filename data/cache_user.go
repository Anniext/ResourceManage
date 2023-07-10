package data

import (
	"ResourceManage/model"
	"sync"
)

type UserMap struct {
	data map[int64]*model.SysBackendUser
	lock sync.RWMutex
}

func NewUserMap() *UserMap {
	return &UserMap{
		data: make(map[int64]*model.SysBackendUser),
	}
}
