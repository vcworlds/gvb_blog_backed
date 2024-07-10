package main

import (
	"gvb_blog/core"
	"gvb_blog/global"
)

func main() {
	// 读取配置文件
	core.InitConfig()
	// 配置gorm
	global.DB = core.InitGorm()
}
