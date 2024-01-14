package config

import "strconv"

type Mysql struct {
	Log_level string `yaml:"log_level"` // 日志等级
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Config    string `yaml:"config"` // 高级配置, 例如 charset
	DB        string `yaml:"db"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
}

func (m Mysql) Dsn() string {
	return m.User + ":" + m.Password +
		"@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" +
		m.DB + m.Config
}
