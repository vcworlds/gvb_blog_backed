package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gvb_blog/config"
)

var (
	Config         *config.Config
	DB             *gorm.DB
	Log            *logrus.Logger
	WhiteImageList = []string{
		"jpg",
		"png",
		"jpeg",
		"ico",
		"tiff",
		"gif",
		"svg",
	}
)
