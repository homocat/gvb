package global

import (
	"main/config"

	"gorm.io/gorm"
)

var (
	Config *config.Confige
	DB *gorm.DB
)