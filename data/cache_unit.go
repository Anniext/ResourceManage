package data

import (
	"ResourceManage/model"
	"ResourceManage/query"
	"log"
	"sync"
)

type UnitMap struct {
	data map[string]*model.AvtUnit
	lock sync.RWMutex
}

func NewUnitMap() *UnitMap {
	return &UnitMap{
		data: make(map[string]*model.AvtUnit),
	}
}

func LoadUnitData() (err error) {
	var unitDataList []*model.AvtUnit
	err = query.AvtUnit.Scan(&unitDataList)
	if err != nil {
		log.Println("avt_unit表数据加载错误：", err)
		return err
	}
	count, _ := query.AvtFile.Count()
	if count > 0 {
		log.Println("avt_unit表缓存数据加载成功!")
		for _, unit := range unitDataList {
			CacheUnit.Set(unit)
		}
	} else {
		log.Println("avt_unit没有数据！")
		return err
	}
	return nil
}

func (m *UnitMap) Set(bu *model.AvtUnit) {
	m.lock.Lock()
	m.data[bu.Name] = bu
	m.lock.Unlock()
}

func (m *UnitMap) Get(name string) *model.AvtUnit {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.data[name]
}

//	func (m *UnitMap) Delete(name string) {
//		m.lock.Lock()
//		defer m.lock.Unlock()
//		if _, ok := m.data[name]; ok {
//			delete(m.data, name)
//		}
//	}
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
	m.data = make(map[string]*model.AvtUnit)
}

func (m *UnitMap) Update(unit *model.AvtUnit) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if _, ok := m.data[unit.Name]; ok {
		delete(m.data, unit.Name)
	}
	m.data[unit.Name] = unit
	_, err := query.AvtUnit.Where(query.AvtUnit.ID.Eq(unit.ID)).Updates(model.AvtUnit{
		Name:       unit.Name,
		Address:    unit.Address,
		UpdateTime: unit.UpdateTime,
	})
	if err != nil {
		log.Println("avt_unit表数据更新错误：", err)
		return
	}
}
