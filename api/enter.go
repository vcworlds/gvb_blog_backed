package api

import "gvb_blog/api/settings_api"

type RouterApp struct {
	SettingsRouter settings_api.SettingsApi
}

var ApiRouterApp = new(RouterApp)
