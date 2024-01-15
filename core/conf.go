package core

import (
	"fmt"
	"os"
	"log"
	"main/config"
	"main/global"

	"gopkg.in/yaml.v2"
)

// InitConf 读取 yaml 文件配置
func InitConf()  {
	const ConfigFile = "settings.yaml"
	c := &config.Confige{}

	yamlConfig, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yaml file err: %s", err))
	}
	err = yaml.Unmarshal(yamlConfig, c)
	if err != nil {
		log.Fatalf("config init unmashal: %s", err)
	}

	log.Println("config load init success")
	global.Config = c
}