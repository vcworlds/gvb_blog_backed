package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb_blog/global"
	"gvb_blog/models"
	"time"
)

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		global.Log.Error("未配置数据库信息")
		return nil
	}
	// 连接数据库
	dsn := global.Config.Mysql.Dsn()
	global.Log.Info("查看数据库连接地址:", dsn)
	var mysqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
		// 开发环境显示的sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}
	//global.MysqlLog = logger.Default.LogMode(logger.Info)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		global.Log.Fatalf(fmt.Sprintf("数据库连接失败: %s", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.ArticleModel{})

	sqlDB.SetMaxIdleConns(10)               // 设置连接池中的最大闲置连接数
	sqlDB.SetMaxOpenConns(100)              // 设置连接池最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 设置连接池最大生存时间，不能超过mysql的wait_timeout
	return db
}
