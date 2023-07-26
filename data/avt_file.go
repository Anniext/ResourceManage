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

type FileList struct {
	Files []model.AvtFile
	Count int64
	Error error
}

func CreateFile(file *model.AvtFile) (bool, string) {
	file.CreateTime = time.Now()
	file.UpdateTime = time.Now()
	file.IsDelete = 0
	file.File = file.Name + "." + file.Type
	id := CacheFile.GetID(file.Name)
	if CacheFile.Get(id) != nil {
		return false, "File name already exists"
	}
	if err := CacheFile.Sync(file); err != nil {
		log.Println("CacheFile Sync err:", err)
		return false, err.Error()
	}
	CacheFile.Set(file)
	return true, ""
}

func GetFile(id int64) (*model.AvtFile, string) {
	if file := CacheFile.Get(id); file != nil {
		return file, ""
	}
	return nil, "File does not exist"
}

func GetFileList(arg *GetHeadBody) interface{} {
	var filelist FileList
	//offset, _ := strconv.Atoi(arg.Offset)
	limit, _ := strconv.Atoi(arg.Limit)
	page, _ := strconv.Atoi(arg.Page)
	v_delete, _ := strconv.ParseInt(arg.Delete, 10, 64)
	if filelist.Error = query.AvtFile.Offset((page * limit) - limit).Limit(limit).Where(query.AvtFile.IsDelete.Eq(v_delete)).Scan(&filelist.Files); filelist.Error != nil {
		return filelist
	}
	filelist.Count, filelist.Error = query.AvtFile.Where(query.AvtFile.IsDelete.Eq(v_delete)).Count()
	if filelist.Error != nil {
		return filelist
	}
	return filelist
}

func UpdateFile(name string, file *model.AvtFile) (bool, string) {
	existingFile, err := GetFile(CacheFile.GetID(name))
	if err != "" {
		return false, err
	}
	existingFile.Name = name
	existingFile.FilePath = file.FilePath
	existingFile.UpdateTime = time.Now()
	existingFile.IsDelete = file.IsDelete
	existingFile.File = name + "." + existingFile.Type
	CacheFile.Update(existingFile)
	return true, ""
}

func DeleteFile(name string) (bool, string) {
	cache := CacheFile.Get(CacheFile.GetID(name))
	if cache == nil {
		return false, "File does not exist"
	}
	if _, err := query.AvtFile.Delete(cache); err != nil {
		return false, err.Error()
	}
	CacheFile.Clear()
	if err := GetFileData(); err != nil {
		return false, "GetFileData err:" + err.Error()
	}
	return true, ""
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
	fileData.File = handler.Filename
	CacheFile.Set(fileData)
	err = CacheFile.Sync(fileData)
	if err != nil {
		return err
	}
	return nil
}
