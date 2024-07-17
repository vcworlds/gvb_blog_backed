package api

import (
	"gvb_blog/api/advert_api"
	"gvb_blog/api/image_api"
	"gvb_blog/api/menu_api"
	"gvb_blog/api/settings_api"
	"gvb_blog/api/user_api"
)

type RouterApp struct {
	SettingsApi settings_api.SettingsApi
	ImageApi    image_api.ImageApi
	AdvertApi   advert_api.AdvertApi
	MenuApi     menu_api.MenuApi
	UserApi     user_api.UserApi
}

var ApiRouterApp = new(RouterApp)
