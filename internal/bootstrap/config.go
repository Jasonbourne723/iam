package bootstrap

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/jasonbourne723/iam/internal/config"
	"github.com/spf13/viper"
)

func InitializeConfig() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		//err =
		panic(fmt.Errorf("read config failed: %s ", err))
	}

	// 监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err := v.Unmarshal(&App.Config); err != nil {
			fmt.Println(err)
		}
	})
	// 将配置赋值给全局变量
	if err := v.Unmarshal(&App.Config); err != nil {
		fmt.Println(err)
	}

	return v
}

type Application struct {
	Config config.Configuration
}

var App = new(Application)
