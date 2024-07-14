package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
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
	return r
}
