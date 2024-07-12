package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	settingGroup := r.Group("settings")
	settingRouter := SettingsRouter{settingGroup}
	settingRouter.SettingsRoute()
	return r
}
