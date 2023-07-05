package data

import "sync"

type UserMap struct {
	data map[int]*AvtUser
	lock sync.RWMutex
}

func NewUserMap() *UserMap {
	return &UserMap{
		data: make(map[int]*AvtUser),
	}
}
