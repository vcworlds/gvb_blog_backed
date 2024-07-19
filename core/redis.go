package core

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gvb_blog/global"
	"time"
)

func ConnectRedis() *redis.Client {
	return ConnectRedisDB(0)
}

func ConnectRedisDB(db int) *redis.Client {
	redisConf := global.Config.Redis
	logrus.Infof("Connecting to Redis at %s", redisConf.Addr()) // 添加日志输出
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr(),   // Redis 服务器地址
		Password: redisConf.Password, // Redis 密码
		DB:       db,                 // 选择的数据库
		PoolSize: redisConf.PoolSize,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		logrus.Errorf("redis连接失败: %v", err)
		return nil
	}
	logrus.Info("Redis连接成功")
	return rdb
}
