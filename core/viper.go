package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go-bbs/global"
	"path/filepath"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		config = "config.yaml"
	} else {
		config = path[0]
	}
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}
	// root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	global.CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v
}
