package main

import (
	"main/core"
	"main/global"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 连接数据库
	global.DB = core.InitGorm()
}