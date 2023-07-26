package data

import (
	"ResourceManage/model"
	"ResourceManage/query"
	"log"
	"sync"
)

type RelaUnitFileMap struct {
	data     map[int64]*model.RelaUnitFile
	dataName map[string]int64
	lock     sync.RWMutex
}

func NewRelaUnitFileMap() *RelaUnitFileMap {
	return &RelaUnitFileMap{
		data:     make(map[int64]*model.RelaUnitFile),
		dataName: make(map[string]int64),
	}
}

func LoadRelaUnitFileData() (err error) {
	var relaUnitFileList []*model.RelaUnitFile
	err = query.RelaUnitFile.Scan(&relaUnitFileList)
	if err != nil {
		log.Println("rela_unit_file Table Data Load Error：", err)
		return err
	}
	count, _ := query.RelaUnitFile.Count()
	if count > 0 {
		log.Println("rela_unit_file Table Data Load Successful!")
		for _, unitFile := range relaUnitFileList {
			CacheRelaUnitFile.Set(unitFile)
		}
	} else {
		log.Println("rela_unit_file Table not Data！")
		return err
	}
	return nil
}
func (m *RelaUnitFileMap) Set(bu *model.RelaUnitFile) {
	m.lock.Lock()
	//m.data[bu.ID] = bu
	m.lock.Unlock()
}

func (m *RelaUnitFileMap) Get(id int64) *model.RelaUnitFile {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.data[id]
}

func (m *RelaUnitFileMap) GetID(name string) int64 {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.dataName[name]
}

func (m *RelaUnitFileMap) Sync(rela *model.RelaUnitFile) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	if err := query.RelaUnitFile.Create(rela); err != nil {
		log.Println("rela_unit_file表数据同步错误：", err)
		return err
	}
	return nil
}
