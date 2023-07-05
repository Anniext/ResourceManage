package data

import "sync"

type UnitMap struct {
	data map[int]*AvtUnit
	lock sync.RWMutex
}

func NewUnitMap() *UnitMap {
	return &UnitMap{
		data: make(map[int]*AvtUnit),
	}
}
