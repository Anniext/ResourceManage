package data

import (
	"ResourceManage/model"
	"ResourceManage/query"
	"log"
	"sync"
)

type UnitMap struct {
	dataName map[string]int64
	data     map[int64]*model.AvtUnit
	lock     sync.RWMutex
}

func NewUnitMap() *UnitMap {
	return &UnitMap{
		data:     make(map[int64]*model.AvtUnit),
		dataName: make(map[string]int64),
	}
}

func LoadUnitData() error {
	unitDataList, err := query.AvtUnit.Preload(query.AvtUnit.FileList, query.AvtUnit.SubUnitList,
		query.AvtUnit.UserList).Find()
	if err != nil {
		log.Println("avt_unit Table Data Load Error：", err)
		return err
	}
	count, _ := query.AvtFile.Count()
	if count > 0 {
		log.Println("avt_unit Table Data Load Successful!")
		for _, unit := range unitDataList {
			CacheUnit.Set(unit)
		}
	} else {
		log.Println("avt_unit Table not Data！")
		return err
	}
	return nil
}

func (m *UnitMap) Set(bu *model.AvtUnit) {
	m.lock.Lock()
	m.data[bu.ID] = bu
	m.dataName[bu.Name] = bu.ID
	m.lock.Unlock()
}

func (m *UnitMap) Get(id int64) *model.AvtUnit {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.data[id]
}

func (m *UnitMap) GetID(name string) int64 {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.dataName[name]
}

func (m *UnitMap) Delete(id int64) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if _, ok := m.data[id]; ok {
		delete(m.data, id)
	}
}
func (m *UnitMap) Sync(unit *model.AvtUnit) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	if err := query.AvtUnit.Create(unit); err != nil {
		log.Println("avt_unit表数据同步错误：", err)
		return err
	}
	return nil
}

func (m *UnitMap) Clear() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.data = make(map[int64]*model.AvtUnit)
	m.dataName = make(map[string]int64)
}

func (m *UnitMap) Update(unit *model.AvtUnit) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if _, ok := m.data[unit.ID]; ok {
		delete(m.data, unit.ID)
		delete(m.dataName, unit.Name)
	}
	_, err := query.AvtUnit.Where(query.AvtUnit.ID.Eq(unit.ID)).Updates(map[string]interface{}{
		"name":        unit.Name,
		"address":     unit.Address,
		"update_time": unit.UpdateTime,
	})
	if err != nil {
		log.Println("avt_unit表数据更新错误：", err)
		return
	}
	newUnit, _ := query.AvtUnit.Where(query.AvtUnit.ID.Eq(unit.ID)).Preload(query.AvtUnit.FileList,
		query.AvtUnit.SubUnitList, query.AvtUnit.UserList).First()
	m.data[unit.ID] = newUnit
}
