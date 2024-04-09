package main

import (
	"calligraphy/apps/activity/rmq/internal/config"
	"calligraphy/apps/activity/rmq/internal/service"
	"flag"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/grab_point.yaml", "the etc file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	srv := service.NewService(c)
	queue := kq.MustNewQueue(c.Kafka, kq.WithHandle(srv.Consume))
	defer queue.Stop()

	fmt.Println("seckill started!!!")
	queue.Start()
}
