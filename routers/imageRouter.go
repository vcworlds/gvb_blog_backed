package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_blog/api/image_api"
)

type ImageRouter struct {
	*gin.Engine
}

func (r ImageRouter) ImageRouter() {
	ImageApi := image_api.NewImageApi()
	r.POST("image", ImageApi.Create)
	r.GET("imageList", ImageApi.Show)
	r.DELETE("delete", ImageApi.Delete)
}
