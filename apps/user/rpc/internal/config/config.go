package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	//Mysql struct {
	//	DataSource string
	//}
	MySQL struct {
		Host        string
		Port        int
		User        string
		Password    string
		Database    string
		TablePrefix string
	}
	CacheRedis cache.CacheConf
	Salt       string
}
