package svc

import (
	"calligraphy/apps/community/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/cache"
)

type ServiceContext struct {
	Config config.Config
	Mysql  struct {
		DataSource string
	}

	CacheRedis cache.CacheConf
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}

}
