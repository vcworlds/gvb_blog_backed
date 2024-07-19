package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_blog/api"
	"gvb_blog/middleware"
)

type UserGroup struct {
	*gin.RouterGroup
}

func (r *UserGroup) UserRouter() {
	userApi := api.ApiRouterApp.UserApi
	r.POST("/register", userApi.Register)
	r.POST("/login", userApi.Login)
	r.GET("/userInfo", middleware.AuthMiddleware(), userApi.UserList)
}
