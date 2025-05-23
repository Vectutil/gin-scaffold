package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Cfg = &Config{}

func InitConfig(filePath string) {
	// 读取配置文件
	if filePath == "" {
		viper.SetConfigFile("config.yaml")
	} else {
		viper.SetConfigFile(filePath)
	}

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return
	}

	err = viper.Unmarshal(Cfg)
	if err != nil {
		fmt.Printf("Error unmarshaling config: %v\n", err)
		return
	}
	fmt.Println(Cfg)
}
