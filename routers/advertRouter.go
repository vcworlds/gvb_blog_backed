package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_blog/api/advert_api"
)

type AdvertRouter struct {
	*gin.RouterGroup
}

func (r AdvertRouter) AdvertRouter() {
	ar := advert_api.NewAdvertApi()
	r.POST("create", ar.Create)
	r.DELETE("delete", ar.Delete)
	r.PUT("update/:id", ar.Update)
	r.GET("show", ar.Show)
}
