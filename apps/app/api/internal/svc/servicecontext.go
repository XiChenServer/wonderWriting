package svc

import (
	"calligraphy/apps/activity/rpc/activityclient"
	"calligraphy/apps/app/api/internal/config"
	"calligraphy/apps/community/rpc/communityclient"

	"calligraphy/apps/grow/rpc/growclient"
	"calligraphy/apps/home/rpc/homeclient"
	"calligraphy/apps/user/rpc/userclient"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"net/http"
)

type ServiceContext struct {
	Config       config.Config
	UserRpc      userclient.User
	CommunityRpc communityclient.Community
	HomeRpc      homeclient.Home
	GrowRpc      growclient.Grow
	Activity     activityclient.Activity
	//KqPusherClient *kq.Pusher
	//ActivityRpc    activityclient.Activity
	GreetMiddleware1 rest.Middleware
	GreetMiddleware2 rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		UserRpc:      userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		CommunityRpc: communityclient.NewCommunity(zrpc.MustNewClient(c.CommunityRpc)),
		HomeRpc:      homeclient.NewHome(zrpc.MustNewClient(c.HomeRpc)),
		GrowRpc:      growclient.NewGrow(zrpc.MustNewClient(c.GroupRpc)),
		Activity:     activityclient.NewActivity(zrpc.MustNewClient(c.ActivityRpc)),
		//KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
		//ActivityRpc:    activityclient.NewActivity(zrpc.MustNewClient(c.ActivityRpc)),
		GreetMiddleware1: greetMiddleware1,
		GreetMiddleware2: greetMiddleware2,
	}
}

func greetMiddleware1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("greetMiddleware1 request ... ")
		next(w, r)
		logx.Info("greetMiddleware1 reponse ... ")
	}
}

func greetMiddleware2(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Info("greetMiddleware2 request ... ")
		next(w, r)
		logx.Info("greetMiddleware2 reponse ... ")
	}
}
