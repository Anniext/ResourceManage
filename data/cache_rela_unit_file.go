package data

import (
	"ResourceManage/model"
	"ResourceManage/query"
	"log"
	"sync"
)

type RelaUnitFileMap struct {
	data map[int64]*model.RelaUnitFile
	lock sync.RWMutex
}

func NewRelaUnitFileMap() *RelaUnitFileMap {
	return &RelaUnitFileMap{
		data: make(map[int64]*model.RelaUnitFile),
	}
}

func LoadRelaUnitFileData() (err error) {
	var relaUnitFileList []*model.RelaUnitFile
	err = query.RelaUnitFile.Scan(&relaUnitFileList)
	if err != nil {
		log.Println("rela_unit_file表数据加载错误：", err)
		return err
	}
	count, _ := query.RelaUnitFile.Count()
	if count > 0 {
		log.Println("rela_unit_file表缓存数据加载成功!")
		for _, unitFile := range relaUnitFileList {
			CacheRelaUnitFile.Set(unitFile)
		}
	} else {
		log.Println("rela_unit_file没有数据！")
		return err
	}
	return nil
}
func (m *RelaUnitFileMap) Set(bu *model.RelaUnitFile) {
	m.lock.Lock()
	m.data[bu.UnitID] = bu
	m.lock.Unlock()
}

func (m *RelaUnitFileMap) Get(unitID int64) *model.RelaUnitFile {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.data[unitID]
}
