package api

import (
	"gvb_blog/api/advert_api"
	"gvb_blog/api/image_api"
	"gvb_blog/api/settings_api"
)

type RouterApp struct {
	SettingsRouter settings_api.SettingsApi
	ImageRouter    image_api.ImageApi
	AdvertRouter   advert_api.AdvertApi
}

var ApiRouterApp = new(RouterApp)
