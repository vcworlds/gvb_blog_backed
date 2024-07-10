package main

import (
	"gvb_blog/core"
	"gvb_blog/global"
	"gvb_blog/routers"
)

func main() {
	// 读取配置文件
	core.InitConfig()
	// 配置日志
	global.Log = core.InitLogger()
	// 配置gorm
	global.DB = core.InitGorm()
	// 路由配置
	router := routers.InitRouter()
	add := global.Config.System.Addr()
	global.Log.Infof("您的程序运行在：%s", add)
	router.Run(add)
}
