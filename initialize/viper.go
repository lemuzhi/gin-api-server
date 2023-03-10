package initialize

import (
	"fmt"
	"gin-api-server/conf"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	vp *viper.Viper
}

func InitConfig(file string) *Config {
	vp := viper.New()
	vp.SetConfigFile(file)
	vp.SetConfigType("toml") // 配置文件的类型
	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error conf file: %s \n", err))
	}
	//配置热更新
	//viper.WatchConfig()
	//viper.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("Config file changed:", e.Name)
	//})

	for k, v := range conf.Configs() {
		readConfig(vp, k, v)
	}

	return &Config{vp: vp}
}

func readConfig(vp *viper.Viper, k string, v interface{}) {
	err := vp.UnmarshalKey(k, v)
	if err != nil {
		log.Panic("Read config file error", err)
	}
}
