package data

import (
	"gorm.io/gorm"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"
	"time"
)

type AvtFile struct {
	Id         int    `gorm:"column:id;auto" json:"id"`
	Name       string `gorm:"column:name;not null" json:"name"`
	Size       int    `gorm:"column:size;not null" json:"size"`
	Type       string `gorm:"column:type;not null" json:"type"`
	FilePath   string `gorm:"column:file_path;not null" json:"file_path"`
	CreateTime string `gorm:"column:create_time;not null" json:"create_time"`
	UpdateTime string `gorm:"column:update_time;not null" json:"update_time"`
	IsDelete   int    `gorm:"column:is_delete;not null" json:"is_delete"`
	UnitId     int    `gorm:"column:unit_id" json:"unit_id"`
	Status     int    `gorm:"column:status;not null" json:"status"`
	File       string `gorm:"column:file;not null" json:"file"`
}

func LoadFileData(db *gorm.DB) (fileDataList []*AvtFile) {
	err := db.Raw("SELECT * FROM avt_file ORDER BY id").Scan(&fileDataList).Error
	if err != nil {
		log.Println("avt_file表数据加载错误：", err)
	}
	count := int64(len(fileDataList))
	if count > 0 {
		log.Println("avt_file表缓存数据加载成功!")
		for _, file := range fileDataList {
			CacheFile.Set(file)
		}
	} else {
		log.Println("没有数据！")
	}
	return
}

func (m *FileMap) Set(bu *AvtFile) {
	m.lock.Lock()
	m.data[bu.Id] = bu
	m.lock.Unlock()
}

func CreateFile(file *AvtFile, db *gorm.DB) error {
	file.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	file.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	file.IsDelete = 0
	file.Status = 1
	file.File = file.Name + "." + file.Type
	if err := db.Table("avt_file").Create(file).Error; err != nil {
		return err
	}
	return nil
}

func GetFile(id string, db *gorm.DB) (*AvtFile, error) {
	var file AvtFile
	if err := db.Table("avt_file").Where("id = ?", id).First(&file).Error; err != nil {
		return nil, err
	}
	return &file, nil
}

func GetFileList(db *gorm.DB, delete_id string) ([]AvtFile, error) {
	var files []AvtFile
	if delete_id == "0" {
		if err := db.Table("avt_file").Find(&files).Error; err != nil {
			return nil, err
		}
	} else if delete_id == "1" {
		if err := db.Table("avt_file").Where("is_delete = ?", 1).Find(&files).Error; err != nil {
			return nil, err
		}
	} else if delete_id == "2" {
		if err := db.Table("avt_file").Where("is_delete = ?", 0).Find(&files).Error; err != nil {
			return nil, err
		}
	}
	return files, nil
}

func UpdateFile(id string, file *AvtFile, db *gorm.DB) error {
	existingFile, err := GetFile(id, db)
	if err != nil {
		log.Println("GetFile err:", err)
		return err
	}
	existingFile.Name = file.Name
	existingFile.FilePath = file.FilePath
	existingFile.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	existingFile.IsDelete = file.IsDelete
	existingFile.Status = file.Status
	existingFile.UnitId = file.UnitId
	existingFile.File = file.Name + "." + existingFile.Type
	if err := db.Table("avt_file").Save(existingFile).Error; err != nil {
		return err
	}

	return nil
}

func DeleteFile(id string, db *gorm.DB) error {
	if err := db.Table("avt_file").Where("id = ?", id).Delete(&AvtFile{}).Error; err != nil {
		return err
	}
	return nil
}

func UploadFile(handler *multipart.FileHeader, f multipart.File, fileData *AvtFile, db *gorm.DB) error {
	//path, _ := os.Getwd()

	fstr := strings.Split(handler.Filename, ".")
	var ftype string
	var fname string
	fname = fstr[0]
	ftype = handler.Header.Get("Content-Type")
	log.Printf("MIME Header: %+v\n", handler.Header)
	// 创建一个新文件，将上传的文件数据写入其中

	tf, err := os.OpenFile("uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(tf)
	_, _ = io.Copy(tf, f)

	if err != nil {
		log.Println(err)
		return err
	}
	fileData.Name = fname
	fileData.Size = int(handler.Size)
	fileData.FilePath = "uploads/"
	fileData.Type = ftype
	fileData.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	fileData.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	fileData.IsDelete = 0
	fileData.Status = 1
	fileData.File = handler.Filename
	if err := db.Table("avt_file").Create(fileData).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DowloadFile(handler *multipart.FileHeader, f multipart.File, fileData *AvtFile, db *gorm.DB) error {
	//path, _ := os.Getwd()

	fstr := strings.Split(handler.Filename, ".")
	var ftype string
	var fname string
	fname = fstr[0]
	ftype = handler.Header.Get("Content-Type")
	log.Printf("MIME Header: %+v\n", handler.Header)
	// 创建一个新文件，将上传的文件数据写入其中

	tf, err := os.OpenFile("uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(tf)
	_, _ = io.Copy(tf, f)

	if err != nil {
		log.Println(err)
		return err
	}
	fileData.Name = fname
	fileData.Size = int(handler.Size)
	fileData.FilePath = "uploads/"
	fileData.Type = ftype
	fileData.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	fileData.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	fileData.IsDelete = 0
	fileData.Status = 1
	if err := db.Table("avt_file").Create(fileData).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}
