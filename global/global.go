package global

import (
	"gorm.io/gorm"
	"gvb_blog/config"
)

var (
	Config *config.Config
	DB     *gorm.DB
)
