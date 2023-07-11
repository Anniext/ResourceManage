package data

import (
	"ResourceManage/model"
	"ResourceManage/query"
	"gorm.io/gorm"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)

type GetHeadBody struct {
	Page   string `json:"page"`
	Limit  string `json:"limit"`
	Offset string `json:"offset"`
	Delete string `json:"delete"`
}

func CreateFile(file *model.AvtFile) string {
	file.CreateTime = time.Now()
	file.UpdateTime = time.Now()
	file.IsDelete = 0
	file.Status = 0
	file.File = file.Name + "." + file.Type
	if CacheFile.Get(file.Name) != nil {
		return "File name already exists"
	}
	CacheFile.Set(file)
	if err := CacheFile.Sync(file); err != nil {
		log.Println("CacheFile Sync err:", err)
		return err.Error()
	}
	return ""
}

func GetFile(name string) (*model.AvtFile, string) {
	if file := CacheFile.Get(name); file != nil {
		return file, ""
	}
	return nil, "File does not exist"
}

func GetFileList(arg *GetHeadBody) ([]model.AvtFile, error) {
	var files []model.AvtFile
	//offset, _ := strconv.Atoi(arg.Offset)
	limit, _ := strconv.Atoi(arg.Limit)
	page, _ := strconv.Atoi(arg.Page)
	v_delete, _ := strconv.ParseInt(arg.Delete, 10, 64)
	err := query.AvtFile.Offset((page * limit) - limit).Limit(limit).Where(query.AvtFile.IsDelete.Eq(v_delete)).Scan(&files)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func UpdateFile(name string, file *model.AvtFile) string {
	existingFile, err := GetFile(name)
	if err != "" {
		return "File does not exist"
	}
	existingFile.Name = file.Name
	existingFile.FilePath = file.FilePath
	existingFile.UpdateTime = time.Now()
	existingFile.IsDelete = file.IsDelete
	existingFile.Status = file.Status
	existingFile.File = file.Name + "." + existingFile.Type
	CacheFile.Update(existingFile)
	return ""
}

func DeleteFile(name string) string {
	if cache := CacheFile.Get(name); cache == nil {
		return "File does not exist"
	}
	file := CacheFile.Get(name)
	if _, err := query.AvtFile.Delete(file); err != nil {
		log.Println("avt_file表数据同步错误：", err)
		return err.Error()
	}
	CacheFile.Clear()
	if err := GetFileData(); err != nil {
		return "GetFileData err:" + err.Error()
	}
	return ""
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
	fileData.Size = handler.Size
	fileData.FilePath = "uploads/"
	fileData.Type = ftype
	fileData.CreateTime = time.Now()
	fileData.UpdateTime = time.Now()
	fileData.IsDelete = 0
	fileData.Status = 1
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
	fileData.IsDelete = 0
	fileData.Status = 1
	if err := db.Table("avt_file").Create(fileData).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}
