package data

import (
	"ResourceManage/model"
	"ResourceManage/query"
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
	if err := CacheFile.Sync(file); err != nil {
		log.Println("CacheFile Sync err:", err)
		return err.Error()
	}
	CacheFile.Set(file)
	return ""
}

func GetFile(name string) (*model.AvtFile, string) {
	if file := CacheFile.Get(name); file != nil {
		return file, ""
	}
	return nil, "File does not exist"
}

func GetFileList(arg *GetHeadBody) ([]model.AvtFile, int64, error) {
	var files []model.AvtFile
	//offset, _ := strconv.Atoi(arg.Offset)
	limit, _ := strconv.Atoi(arg.Limit)
	page, _ := strconv.Atoi(arg.Page)
	v_delete, _ := strconv.ParseInt(arg.Delete, 10, 64)
	if err := query.AvtFile.Offset((page * limit) - limit).Limit(limit).Where(query.AvtFile.IsDelete.Eq(v_delete)).Scan(&files); err != nil {
		return nil, 0, err
	}
	count, err := query.AvtFile.Where(query.AvtFile.IsDelete.Eq(v_delete)).Count()
	if err != nil {
		return nil, 0, err
	}
	return files, count, nil
}

func UpdateFile(name string, file *model.AvtFile) string {
	existingFile, err := GetFile(name)
	if err != "" {
		return err
	}
	existingFile.Name = name
	existingFile.FilePath = file.FilePath
	existingFile.UpdateTime = time.Now()
	existingFile.IsDelete = file.IsDelete
	existingFile.Status = file.Status
	existingFile.File = name + "." + existingFile.Type
	CacheFile.Update(existingFile)
	return ""
}

func DeleteFile(name string) string {
	cache := CacheFile.Get(name)
	if cache == nil {
		return "File does not exist"
	}
	if _, err := query.AvtFile.Delete(cache); err != nil {
		log.Println("avt_file表数据同步错误：", err)
		return err.Error()
	}
	CacheFile.Clear()
	if err := GetFileData(); err != nil {
		return "GetFileData err:" + err.Error()
	}
	return ""
}

func UploadFile(handler *multipart.FileHeader, f multipart.File, fileData *model.AvtFile) error {
	fstr := strings.Split(handler.Filename, ".")
	fname := fstr[0]
	ftypelist := strings.Split(handler.Header.Get("Content-Type"), "/")
	ftype := ftypelist[len(ftypelist)-1]
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
	fileData.Status = 0
	fileData.File = handler.Filename
	CacheFile.Set(fileData)
	err = CacheFile.Sync(fileData)
	if err != nil {
		return err
	}
	return nil
}
