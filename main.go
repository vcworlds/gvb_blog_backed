package main

import (
	"gvb_blog/core"
	_ "gvb_blog/docs"
	"gvb_blog/flag"
	"gvb_blog/global"
	"gvb_blog/routers"
)

// @title           gvb_API
// @version         1.0
// @description     API 文档
// @host     127.0.0.1：8080
// @BasePath  /
func main() {
	// 读取配置文件
	core.InitConfig()
	// 配置日志
	global.Log = core.InitLogger()
	// 配置gorm
	global.DB = core.InitGorm()
	// 配置redis
	global.Redis = core.ConnectRedis()
	if global.Redis == nil {
		global.Log.Fatal("Redis client initialization failed")
		return
	}
	// 命令行迁移
	option := flag.Parse()
	if flag.IsStopWeb(&option) {
		flag.SwitchOption(&option)
		return
	}
	// 路由配置
	r := routers.InitRouter()
	add := global.Config.System.Addr()
	global.Log.Infof("您的程序运行在：%s", add)
	err := r.Run(add)
	if err != nil {
		global.Log.Infof("程序运行失败", err)
	}
}
