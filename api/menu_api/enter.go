package menu_api

import (
	"gvb_blog/response"
)

type IMenuApi interface {
	response.RestApi
}

type MenuApi struct {
}

func NewMenuApi() IMenuApi {
	return MenuApi{}
}
