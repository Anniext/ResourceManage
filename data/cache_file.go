package data

import (
	"ResourceManage/model"
	"ResourceManage/query"
	"log"
	"sync"
)

type FileMap struct {
	data     map[int64]*model.AvtFile
	lock     sync.RWMutex
	dataName map[string]int64
}

func NewFileMap() *FileMap {
	return &FileMap{
		data:     make(map[int64]*model.AvtFile),
		dataName: make(map[string]int64),
	}
}

func LoadFileData() error {
	fileDataList, err := query.AvtFile.Preload(query.AvtFile.UnitList).Find()
	if err != nil {
		log.Println("avt_file Table Data Load Error：", err)
		return err
	}
	count, _ := query.AvtFile.Count()
	if count > 0 {
		log.Println("avt_file Table Data Load Successful!")
		for _, file := range fileDataList {
			CacheFile.Set(file)
		}
	} else {
		log.Println("avt_file Table not Data！")
		return err
	}
	return nil
}

func (m *FileMap) Set(bu *model.AvtFile) {
	m.lock.Lock()
	m.data[bu.ID] = bu
	m.dataName[bu.Name] = bu.ID
	m.lock.Unlock()
}

func (m *FileMap) Get(id int64) *model.AvtFile {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.data[id]
}

func (m *FileMap) GetID(name string) int64 {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.dataName[name]
}

func (m *FileMap) Delete(id int64) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if _, ok := m.data[id]; ok {
		delete(m.data, id)
	}
}

func (m *FileMap) Sync(file *model.AvtFile) error {
	m.lock.Lock()
	defer m.lock.Unlock()
	if err := query.AvtFile.Create(file); err != nil {
		log.Println("avt_file表数据同步错误：", err)
		return err
	}
	return nil
}

func (m *FileMap) Clear() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.data = make(map[int64]*model.AvtFile)
	m.dataName = make(map[string]int64)
}

func (m *FileMap) Update(file *model.AvtFile) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if _, ok := m.data[file.ID]; ok {
		delete(m.data, file.ID)
		delete(m.dataName, file.Name)
		m.data[file.ID] = file
		m.dataName[file.Name] = file.ID
	}
	m.data[file.ID] = file
	_, err := query.AvtFile.Where(query.AvtFile.ID.Eq(file.ID)).Updates(map[string]interface{}{
		"name":        file.Name,
		"type":        file.Type,
		"file":        file.File,
		"update_time": file.UpdateTime,
		"is_delete":   file.IsDelete,
	})
	if err != nil {
		log.Println("avt_file表数据更新错误：", err)
		return
	}
}
