package data

import (
	"ResourceManage/dao"
	"ResourceManage/model"
	"log"
	"sync"
)

type FileMap struct {
	data map[int64]*model.AvtFile
	lock sync.RWMutex
}

func NewFileMap() *FileMap {
	return &FileMap{
		data: make(map[int64]*model.AvtFile),
	}
}

func LoadFileData() (err error) {
	var fileDataList []*model.AvtFile
	err = dao.AvtFile.Scan(&fileDataList)
	if err != nil {
		log.Println("avt_file表数据加载错误：", err)
		return err
	}
	count, _ := dao.AvtFile.Count()
	if count > 0 {
		log.Println("avt_file表缓存数据加载成功!")
		for _, file := range fileDataList {
			CacheFile.Set(file)
		}
	} else {
		log.Println("avt_file没有数据！")
		return err
	}
	return nil
}

func (m *FileMap) Set(bu *model.AvtFile) {
	m.lock.Lock()
	m.data[bu.ID] = bu
	m.lock.Unlock()
}
