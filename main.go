package main

import (
	"gvb_blog/core"
	"gvb_blog/flag"
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
