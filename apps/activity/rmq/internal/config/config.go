package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	Kafka   kq.KqConf
	UserRpc zrpc.RpcClientConf
	MySQL   struct {
		Host        string
		Port        int
		User        string
		Password    string
		Database    string
		TablePrefix string
	}
	ActivityRpc zrpc.RpcClientConf
}
