package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"main/config"

	"gopkg.in/yaml.v2"
)

// InitConf 读取 yaml 文件配置
func InitConf()  {
	const ConfigFile = "settings.yaml"
	c := &config.Confige{}

	yamlConfig, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yaml file err: %s", err))
	}
	err = yaml.Unmarshal(yamlConfig, c)
	if err != nil {
		log.Fatalf("config init unmashal: %s", err)
	}

	log.Println("config load init success")
	fmt.Println(c)
}