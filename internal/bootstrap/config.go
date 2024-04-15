package bootstrap

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/jasonbourne723/iam/internal/global"
	"github.com/spf13/viper"
)

func InitializeConfig() *viper.Viper {

	v := viper.New()
	v.SetConfigFile("../../../config.yaml")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s ", err))
	}

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		if err := v.Unmarshal(&global.App.Config); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.App.Config); err != nil {
		fmt.Println(err)
	}
	return v
}
