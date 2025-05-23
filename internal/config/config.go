package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var Cfg = &Config{}

func InitConfig() {
	// 读取配置文件
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	// 解析 YAML 数据
	err = yaml.Unmarshal(file, Cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(Cfg)
}
