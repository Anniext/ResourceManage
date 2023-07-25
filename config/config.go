package config

import (
	"github.com/bytedance/sonic"
	"log"
	"os"
)

type ConfPath struct {
	Path string
}

type ConfBody struct {
	Mode string `json:"mode"`
	Dev  DevBody    `json:"dev"`
	Pub  PubBody    `json:"pub"`
}

type DevBody struct {
	AppName string `json:"app_name"`
	DSN     string `json:"dsn"`
	Router  Router `json:"router"`
	Target  Target `json:"target"`
}


type PubBody struct {
	AppName string `json:"app_name"`
	DSN     string `json:"dsn"`
	Router  Router `json:"router"`
	Target  Target `json:"target"`
	Mqtt    *Mqtt  `json:"mqtt,omitempty"`
}
type Mqtt struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Router struct {
	Host string `json:"host"`
	Mode string `json:"mode"`
}

type Target struct {
	Download string `json:"download"`
	Upload   string `json:"upload"`
}

var (
	Configs ConfBody
	Config  ConfigsInterface
)

func (c *ConfPath) Read() {
	file, err := os.ReadFile(c.Path)
	if err != nil {
		log.Println("Failed to read configuration file:", err)
	} else {
		log.Println("Configuration file loaded successfully...")
	}
	err = sonic.Unmarshal(file, &Configs)
	if err != nil {
		log.Println("Failed to serialize json object:", err)
	} else {
		log.Println("serialized json object successfully...")
		log.Println("<----------------------------->")
	}
}
