package initialize

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig(file string) {
	viper.New()
	viper.SetConfigFile(file)
	viper.SetConfigType("toml") // 配置文件的类型
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//配置热更新
	//viper.WatchConfig()
	//viper.OnConfigChange(func(e fsnotify.Event) {
	//	fmt.Println("Config file changed:", e.Name)
	//})
}
