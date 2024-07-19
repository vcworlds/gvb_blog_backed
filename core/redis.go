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
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr(),   // Redis 服务器地址
		Password: redisConf.Password, // Redis 密码
		DB:       0,                  // 选择的数据库
		PoolSize: redisConf.PoolSize,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		logrus.Errorf("redis连接失败%s", redisConf.Addr())
		return nil
	}
	return rdb
}
