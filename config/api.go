package config

func init() {
	Config = NewConfig()
	Config.Read()
	Log = NewLog()
}

func Configinit() {
	Log.Write()
}

type ConfigsInterface interface {
	Read()
}

func NewConfig() *ConfPath {
	return &ConfPath{Path: "./cnf/config.json"}
}

func NewLog() *LogPath {
	return &LogPath{Path: "./log/log.txt"}
}
