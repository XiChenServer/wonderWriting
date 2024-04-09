package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	UserRpc      zrpc.RpcClientConf
	CommunityRpc zrpc.RpcClientConf
	HomeRpc      zrpc.RpcClientConf
	GroupRpc     zrpc.RpcClientConf
	//KqPusherConf struct {
	//	Brokers []string
	//	Topic   string
	//}
	//KqConsumerConf kq.KqConf
	ActivityRpc zrpc.RpcClientConf
}

//
//docker run -d --name kafka -p 9092:9092 -e KAFKA_BROKER_ID=0 -e KAFKA_ZOOKEEPER_CONNECT=127.0.0.1:2181 -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://127.0.0.1:9092 -e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092 wurstmeister/kafka
