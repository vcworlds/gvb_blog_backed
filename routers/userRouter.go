package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_blog/api"
)

type UserGroup struct {
	*gin.RouterGroup
}

func (r *UserGroup) UserRouter() {
	var userApi = api.ApiRouterApp
	r.POST("/register", userApi.UserApi.Register)
	r.POST("/login", userApi.UserApi.Login)
}
