package config

import (
	"log"
	"os"
)

type LogPath struct {
	Path string
}

type LogConfBody struct {
	Info    *log.Logger
	Error   *log.Logger
	Warning *log.Logger
}

type LogsInterface interface {
	Write()
	Detection()
}

var (
	Logs LogConfBody
	Log  LogsInterface
)

func (c *LogPath) Write() {
	file, _ := os.Create(c.Path) //需要运行读取配置文件函数之后
	Logs.Info = log.New(file, "[Info]", log.Ldate|log.Ltime|log.Lshortfile)
	Logs.Error = log.New(file, "[Error]", log.Ldate|log.Ltime|log.Lshortfile)
	Logs.Warning = log.New(file, "[Warning]", log.Ldate|log.Ltime|log.Lshortfile)
	log.Println("Log file created successfully, program starts running...")
	log.Println("<----------------------------->")
}

func (c *LogPath) Detection() {
	maxSize := 1024 * 1024 * 10 // 10M
	fileInfo, err := os.Stat(c.Path)
	if err != nil {
		Logs.Error.Println("Failed to get log file information:", err)
	}
	if fileInfo.Size() > int64(maxSize) {
		err = os.Rename(c.Path, c.Path+".bak")
		if err != nil {
			Logs.Error.Println("Failed to rename log file:", err)
		}
		_, err := os.Create(c.Path)
		Logs.Info.Println("Log file too large, recreate new log file")
		if err != nil {
			Logs.Error.Println("Failed to create new log file:", err)
		}
	} else {
		Logs.Info.Println("Log file size is normal...")
	}
}
