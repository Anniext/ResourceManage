package data

import (
	"ResourceManage/model"
	"gorm.io/gorm"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"
	"time"
)

func CreateFile(file *model.AvtFile, db *gorm.DB) error {
	file.CreateTime = time.Now()
	file.UpdateTime = time.Now()
	file.IsDelete = "0"
	file.Status = "1"
	file.File = file.Name + "." + file.Type
	if err := db.Table("avt_file").Create(file).Error; err != nil {
		return err
	}
	return nil
}

func GetFile(id string, db *gorm.DB) (*model.AvtFile, error) {
	var file model.AvtFile
	if err := db.Table("avt_file").Where("id = ?", id).First(&file).Error; err != nil {
		return nil, err
	}
	return &file, nil
}

func GetFileList(db *gorm.DB, delete_id string) ([]model.AvtFile, error) {
	var files []model.AvtFile
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

func UpdateFile(id string, file *model.AvtFile, db *gorm.DB) error {
	existingFile, err := GetFile(id, db)
	if err != nil {
		log.Println("GetFile err:", err)
		return err
	}
	existingFile.Name = file.Name
	existingFile.FilePath = file.FilePath
	existingFile.UpdateTime = time.Now()
	existingFile.IsDelete = file.IsDelete
	existingFile.Status = file.Status
	existingFile.UnitID = file.UnitID
	existingFile.File = file.Name + "." + existingFile.Type
	if err := db.Table("avt_file").Save(existingFile).Error; err != nil {
		return err
	}

	return nil
}

func DeleteFile(id string, db *gorm.DB) error {
	if err := db.Table("avt_file").Where("id = ?", id).Delete(&model.AvtFile{}).Error; err != nil {
		return err
	}
	return nil
}

func UploadFile(handler *multipart.FileHeader, f multipart.File, fileData *model.AvtFile, db *gorm.DB) error {
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
	fileData.Size = int64(handler.Size)
	fileData.FilePath = "uploads/"
	fileData.Type = ftype
	fileData.CreateTime = time.Now()
	fileData.UpdateTime = time.Now()
	fileData.IsDelete = "0"
	fileData.Status = "1"
	fileData.File = handler.Filename
	if err := db.Table("avt_file").Create(fileData).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DowloadFile(handler *multipart.FileHeader, f multipart.File, fileData *model.AvtFile, db *gorm.DB) error {
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
	fileData.Size = int64(int(handler.Size))
	fileData.FilePath = "uploads/"
	fileData.Type = ftype
	fileData.CreateTime = time.Now()
	fileData.UpdateTime = time.Now()
	fileData.IsDelete = "0"
	fileData.Status = "1"
	if err := db.Table("avt_file").Create(fileData).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}
