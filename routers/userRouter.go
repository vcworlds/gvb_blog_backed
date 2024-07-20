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
	r.PUT("/userRole", middleware.AuthMiddlewareAdmin(), userApi.RoleUpdate)
	r.PUT("/userPassword", middleware.AuthMiddleware(), userApi.PasswordUpdate)
	r.GET("/userLogout", middleware.AuthMiddleware(), userApi.UserLogout)
	r.DELETE("/userDelete", middleware.AuthMiddlewareAdmin(), userApi.UserDelete)
}
