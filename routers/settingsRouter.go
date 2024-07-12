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
	r.GET("/settingInfo", settingsApi.SettingsRouter.SettingsInfo)
	r.PUT("/updateSetting", settingsApi.SettingsRouter.SettingsUpdate)
}
