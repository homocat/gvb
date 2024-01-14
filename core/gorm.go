package core

import (
	"fmt"
	"log"
	"main/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		log.Printf("未配置mysql, 取消连接gorm")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()

	var mysqlLoger logger.Interface
	if global.Config.System.Env == "debug" {
		// 开发环境显示所有sql
		mysqlLoger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLoger = logger.Default.LogMode(logger.Error) // 只打印错误的sql
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLoger,
	})
	if err != nil {
		log.Fatalf(fmt.Sprintf("[%s] mysql连接失败", dsn))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接
	sqlDB.SetMaxOpenConns(100)              // 最大可连接
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间, 不能超过mysql的wait_timeout

	return db
}