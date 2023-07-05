package data

import "sync"

type FileMap struct {
	data map[int]*AvtFile
	lock sync.RWMutex
}

func NewFileMap() *FileMap {
	return &FileMap{
		data: make(map[int]*AvtFile),
	}
}
