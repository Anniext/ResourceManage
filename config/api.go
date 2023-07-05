package config

func init() {
	Config = NewConfig()
	Config.Read()
	Log = NewLog()

}

func Configinit() {
	Log.Write()
}
