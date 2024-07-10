package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_blog/api"
)

type SettingsRouter struct {
	*gin.Engine
}

func (r SettingsRouter) SettingsRoute() {
	settingsApi := api.RouterApp{}
	r.GET("/", settingsApi.SettingsRouter.SettingsInfoView)
}
