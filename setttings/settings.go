package setttings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init()(err error){
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err=viper.ReadInConfig()
	if err!=nil{
		fmt.Printf("Fatal error config file:%s\n",err.Error())
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event){
		fmt.Println("配置文件被修改")
	})
	return
}
