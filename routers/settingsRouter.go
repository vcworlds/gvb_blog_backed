package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_blog/api"
)

type SettingsRouter struct {
	*gin.RouterGroup
}

func (r SettingsRouter) SettingsRoute() {
	settingsApi := api.ApiRouterApp
	r.GET("/settingInfo/:name", settingsApi.SettingsApi.SettingsSiteInfo)
	r.PUT("/updateSetting/:name", settingsApi.SettingsApi.SettingsSiteInfoUpdate)
}
