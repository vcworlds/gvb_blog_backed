package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_blog/api"
)

type SettingsRouter struct {
	*gin.RouterGroup
}

func (r SettingsRouter) SettingsRoute() {
	settingsApi := api.RouterApp{}
	r.GET("/settingInfo", settingsApi.SettingsRouter.SettingsInfo)
}
