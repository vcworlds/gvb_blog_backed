package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_blog/api"
)

type ImageRouter struct {
	*gin.Engine
}

func (r ImageRouter) ImageRouter() {
	ImageApi := api.ApiRouterApp
	r.POST("image", ImageApi.ImageRouter.ImageView)
	r.GET("imageList", ImageApi.ImageRouter.ImageList)
}
