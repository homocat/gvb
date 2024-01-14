package config

import "fmt"

type System struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

func (sys System) Addr() string {
	return fmt.Sprintf("%s:%d", sys.Host, sys.Port)
}
