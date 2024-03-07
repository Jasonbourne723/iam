package global

import (
	"github.com/jasonbourne723/iam/internal/config"
	"gorm.io/gorm"
)

type Application struct {
	Config *config.Configuration
	DB     *gorm.DB
}

var App = new(Application)
