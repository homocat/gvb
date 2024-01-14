package global

import (
	"main/config"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config *config.Confige
	DB     *gorm.DB
	Log    *logrus.Logger
)
