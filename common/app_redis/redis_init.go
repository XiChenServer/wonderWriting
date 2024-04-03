package app_redis

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"time"
)

var Redis *redis.Redis

func init() {
	conf := redis.RedisConf{
		Host:        "127.0.0.1:6379",
		Type:        "node",
		Pass:        "",
		Tls:         false,
		NonBlock:    false,
		PingTimeout: time.Second,
	}
	Redis = redis.MustNewRedis(conf)
}
