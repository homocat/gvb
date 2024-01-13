package config

type Mysql struct {
	Log_level string `yaml:"log_level"` // 日志等级
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Db        string `yaml:"db"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
}