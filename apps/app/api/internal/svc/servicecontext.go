package svc

import (
	"calligraphy/apps/app/api/internal/config"
	"calligraphy/apps/community/rpc/communityclient"
	"calligraphy/apps/home/rpc/homeclient"
	"calligraphy/apps/user/rpc/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	UserRpc      userclient.User
	CommunityRpc communityclient.Community
	HomeRpc      homeclient.Home
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		UserRpc:      userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		CommunityRpc: communityclient.NewCommunity(zrpc.MustNewClient(c.CommunityRpc)),
		HomeRpc:      homeclient.NewHome(zrpc.MustNewClient(c.HomeRpc)),
	}
}
