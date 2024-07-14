package advert_api

import (
	"gvb_blog/response"
)

type IAdvertApi interface {
	response.RestApi
}

type AdvertApi struct {
}

func NewAdvertApi() IAdvertApi {
	return AdvertApi{}
}
