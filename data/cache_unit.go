package data

import (
	"ResourceManage/dao"
	"ResourceManage/model"
	"log"
	"sync"
)

type UnitMap struct {
	data map[int64]*model.AvtUnit
	lock sync.RWMutex
}

func NewUnitMap() *UnitMap {
	return &UnitMap{
		data: make(map[int64]*model.AvtUnit),
	}
}

func LoadUnitData() (err error) {
	var unitDataList []*model.AvtUnit
	err = dao.AvtFile.Scan(&unitDataList)
	if err != nil {
		log.Println("avt_unit表数据加载错误：", err)
		return err
	}
	count, _ := dao.AvtFile.Count()
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
	m.data[bu.ID] = bu
	m.lock.Unlock()
}
