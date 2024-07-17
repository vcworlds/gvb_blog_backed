package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	settingGroup := r.Group("settings")
	// 系统配置
	settingRouter := SettingsRouter{settingGroup}
	settingRouter.SettingsRoute()
	// 图片配置
	imageRouter := ImageRouter{r}
	imageRouter.ImageRouter()
	// 广告配置
	aGroup := r.Group("advert")
	advertRouter := AdvertRouter{aGroup}
	advertRouter.AdvertRouter()

	// 菜单管理
	mGroup := r.Group("menu")
	menuRouter := MenuRouter{mGroup}
	menuRouter.MenuRouter()
	// 用户管理
	uGroup := r.Group("user")
	userRouter := UserGroup{uGroup}
	userRouter.UserRouter()
	return r
}
