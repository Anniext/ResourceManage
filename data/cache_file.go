package data

import (
	"ResourceManage/model"
	"ResourceManage/query"
	"log"
	"sync"
)

type FileMap struct {
	data map[string]*model.AvtFile
	lock sync.RWMutex
}

func NewFileMap() *FileMap {
	return &FileMap{
		data: make(map[string]*model.AvtFile),
	}
}

func LoadFileData() (err error) {
	var fileDataList []*model.AvtFile
	err = query.AvtFile.Scan(&fileDataList)
	if err != nil {
		log.Println("avt_file表数据加载错误：", err)
		return err
	}
	count, _ := query.AvtFile.Count()
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
	m.data[bu.Name] = bu
	m.lock.Unlock()
}

func (m *FileMap) Get(name string) *model.AvtFile {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.data[name]
}

func (m *FileMap) Delete(name string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if _, ok := m.data[name]; ok {
		delete(m.data, name)
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
	m.data = make(map[string]*model.AvtFile)
}

func (m *FileMap) Update(file *model.AvtFile) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if _, ok := m.data[file.Name]; ok {
		delete(m.data, file.Name)
		m.data[file.Name] = file
	}
	m.data[file.Name] = file
	_, err := query.AvtFile.Where(query.AvtFile.ID.Eq(file.ID)).Updates(map[string]interface{}{
		"name":        file.Name,
		"type":        file.Type,
		"file":        file.File,
		"update_time": file.UpdateTime,
		"is_delete":   file.IsDelete,
		"status":      file.Status,
	})
	if err != nil {
		log.Println("avt_file表数据更新错误：", err)
		return
	}
}
