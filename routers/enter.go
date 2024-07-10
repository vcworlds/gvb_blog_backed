package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	settingsRouter := SettingsRouter{
		r,
	}
	settingsRouter.SettingsRoute()
	return r
}
