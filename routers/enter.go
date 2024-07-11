package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	settingGroup := r.Group("settings")
	settingsRouter := SettingsRouter{
		settingGroup,
	}
	settingsRouter.SettingsRoute()
	return r
}
