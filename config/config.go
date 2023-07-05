package config

import (
	"encoding/json"
	"log"
	"os"
)

type ConfPath struct {
	Path string
}

type ConfBody struct {
	Dsn          string `json:"dsn"`
	AppPort      string `json:"app_port"`
	DownloadPort string `json:"download_port"`
	UploadPort   string `json:"upload_port"`
}

var (
	Configs ConfBody
	Config  ConfigsInterface
)

type ConfigsInterface interface {
	Read()
}

func NewConfig() *ConfPath {
	return &ConfPath{Path: "./cnf/config.json"}
}

// 读取文件的方法
func (c *ConfPath) Read() {
	// 读取json文件
	file, err := os.ReadFile(c.Path)
	if err != nil {
		log.Println("Failed to read configuration file:", err)
	} else {
		log.Println("Configuration file loaded successfully...")
	}
	// 解析json文件,序列化为对象
	err = json.Unmarshal(file, &Configs)
	if err != nil {
		log.Println("Failed to serialize json object:", err)
	} else {
		log.Println("serialized json object successfully...")
		log.Println("<----------------------------->")
	}
}
