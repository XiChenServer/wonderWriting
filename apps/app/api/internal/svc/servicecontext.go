package svc

import (
	"calligraphy/apps/activity/rpc/activityclient"
	"calligraphy/apps/app/api/internal/config"
	"calligraphy/apps/community/rpc/communityclient"
	"calligraphy/apps/group/rpc/groupclient"
	"calligraphy/apps/home/rpc/homeclient"
	"calligraphy/apps/user/rpc/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	UserRpc      userclient.User
	CommunityRpc communityclient.Community
	HomeRpc      homeclient.Home
	GroupRpc     groupclient.Group
	Activity     activityclient.Activity
	//KqPusherClient *kq.Pusher
	//ActivityRpc    activityclient.Activity
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		UserRpc:      userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		CommunityRpc: communityclient.NewCommunity(zrpc.MustNewClient(c.CommunityRpc)),
		HomeRpc:      homeclient.NewHome(zrpc.MustNewClient(c.HomeRpc)),
		GroupRpc:     groupclient.NewGroup(zrpc.MustNewClient(c.GroupRpc)),
		Activity:     activityclient.NewActivity(zrpc.MustNewClient(c.ActivityRpc)),
		//KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
		//ActivityRpc:    activityclient.NewActivity(zrpc.MustNewClient(c.ActivityRpc)),
	}
}
