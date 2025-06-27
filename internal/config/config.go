package config

import (
	"fmt"
	"gin-scaffold/pkg/logger"
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
		logger.Logger.Error(fmt.Sprintf("Error reading config file: %v\n", err))
		return
	}

	err = viper.Unmarshal(Cfg)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("Error unmarshaling config: %v\n", err))
		return
	}
	if Cfg.System.Env == "dev" {
		//indent, _ := json.MarshalIndent(Cfg, "", "  ")
		//fmt.Println(string(indent))
	}
}
