package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_blog/api/image_api"
)

type ImageRouter struct {
	*gin.Engine
}

func (r ImageRouter) ImageRouter() {
	ImageApi := image_api.ImageApi{}
	r.POST("image", ImageApi.ImageView)
}
