package main

import (
	"main/core"
	"main/global"

	"github.com/sirupsen/logrus"
)

func main() {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()
	global.Log.Warn("asfd")
	global.Log.Infof("asdf")
	global.Log.Error("asdf")
logrus.Warn("asf")
logrus.Debug("asdf")
	// 连接数据库
	global.DB = core.InitGorm()
}