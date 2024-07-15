package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_blog/api/menu_api"
)

type MenuRouter struct {
	*gin.RouterGroup
}

func (m MenuRouter) MenuRouter() {
	menuApi := menu_api.NewMenuApi()
	m.POST("create", menuApi.Create)
	m.PUT("update", menuApi.Update)
	m.DELETE("delete", menuApi.Delete)
	m.GET("show", menuApi.Show)
}
