package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"time"

	"calligraphy/apps/activity/rpc/internal/config"
	"calligraphy/apps/activity/rpc/internal/server"
	"calligraphy/apps/activity/rpc/internal/svc"
	"calligraphy/apps/activity/rpc/types/activity"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/activity.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		activity.RegisterActivityServer(grpcServer, server.NewActivityServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	defer s.Stop()
	// 启动定时器
	go startTimer(ctx)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

func startTimer(ctx *svc.ServiceContext) {
	for {
		// 获取当前时间
		now := time.Now()

		// 确定每天更新积分的时间点（例如每天的凌晨零点）
		updateTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
		durationUntilUpdate := updateTime.Sub(now)

		// 设置定时器，在每天的更新时间点执行更新积分的操作
		timer := time.NewTimer(durationUntilUpdate)
		<-timer.C

		// 调用更新积分的逻辑
		success := updatePointsOneDay(ctx)
		if !success {
			//log.Println("Failed to update points.")
		}
	}
}
func updatePointsOneDay(ctx *svc.ServiceContext) bool {
	// 获取当前时间
	now := time.Now()

	// 确定每天更新积分的时间点（例如每天的凌晨零点）
	updateTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	// 判断当前时间是否已经过了更新时间，如果已经过了则表示今天已经更新过积分，无需再次更新
	if now.After(updateTime) {
		return false
	}

	// 将积分数量存储到 Redis 中，并设置过期时间为一天
	count := "10000" // 假设初始积分为 10000
	expiration := 24 * time.Hour
	err := ctx.RDB.Setex("point_count_one_day", count, int(expiration))
	if err != nil {
		logx.Errorf("failed to update points count in Redis: %v", err)
		return false
	}

	// 返回更新成功
	return true
}
