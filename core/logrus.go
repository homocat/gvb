package core

import (
	"bytes"
	"fmt"
	"main/global"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

// 颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

// logrus formatter
type LogFormatter struct{}
// log 配置

// Format 实现 (entry *logrus.Entry) ([]byte, error) 接口
func (t LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	log := global.Config.Logger
	// 根据不同 level 显示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	// 日志缓冲区
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	// 自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		// 自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		// 自定义输出格式
		fmt.Fprintf(b, "%s [%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", 
		log.Prefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "%s [%s] \x1b[%dm[%s]\x1b[0m %s\n", 
		log.Prefix, timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

// InitLogger 用于初始化日志记录器
func InitLogger() *logrus.Logger {
	mLog := logrus.New() // 创建一个新的logrus.Logger实例
	// 将日志输出设置为标准输出
	mLog.SetOutput(os.Stdout)
	// 设置日志的格式化器为LogFormatter
	mLog.SetFormatter(&LogFormatter{})
	// 设置是否返回函数名和行号
	mLog.SetReportCaller(global.Config.Logger.ShowLine)

	// 解析配置文件中的日志级别
	level, err := logrus.ParseLevel(string(global.Config.Logger.Level))
	if err != nil {
		level = logrus.InfoLevel
	}
	// 设置最低的level
	mLog.SetLevel(level)
	// 设置全局log
	InitDefualtLogger()

	// 返回日志记录器的指针
	return mLog
}

func InitDefualtLogger() {
	// 全局log
	logrus.SetOutput(os.Stdout)                           // 设置输出类型
	logrus.SetReportCaller(global.Config.Logger.ShowLine) // 开启返回函数名和行号
	logrus.SetFormatter(LogFormatter{})                   // 设置自己定义的formatter
	logrus.SetLevel(logrus.DebugLevel)
}
