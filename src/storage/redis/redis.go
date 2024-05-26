package redis

import (
	"GuGoTik/src/constant/config"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"strings"
)

var Client redis.UniversalClient

func init() {
	addrs := strings.Split(config.EnvCfg.RedisAddr, ";")
	// Universal 通用的
	Client = redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:      addrs,
		Password:   config.EnvCfg.RedisPassword,
		DB:         config.EnvCfg.RedisDB,
		MasterName: config.EnvCfg.RedisMaster,
	})
	// todo 未设置链路追踪服务商
	if err := redisotel.InstrumentTracing(Client); err != nil {
		panic(err)
	}
	// todo 未设置检测踪服务商
	if err := redisotel.InstrumentMetrics(Client); err != nil {
		panic(err)
	}
}
