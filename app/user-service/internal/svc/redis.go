package svc

import (
	"context"

	"go-zero-boilerplate/app/user-service/internal/config"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
)

func MustNewRedis(redisCfg config.RedisConf) redis.UniversalClient {
	var client redis.UniversalClient
	// 使用集群模式
	if redisCfg.UseCluster {
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    redisCfg.ClusterAddrs,
			Password: redisCfg.Password,
		})
	} else {
		// 单节点
		client = redis.NewClient(&redis.Options{
			Addr:     redisCfg.Addr,
			Password: redisCfg.Password,
			DB:       redisCfg.DB,
		})
	}
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		logx.Error("redis connect ping failed, err:", err)
		panic(err)
	}
	logx.Info("redis connect ping response, pong:", pong)
	return client
}