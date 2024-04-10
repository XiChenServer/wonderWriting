package service

import (
	"calligraphy/apps/activity/rmq/internal/config"
	"calligraphy/apps/user/rpc/userclient"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"log"
	"sync"
)

const (
	chanCount   = 10
	bufferCount = 1024
)

type Service struct {
	c       config.Config
	UserRpc userclient.User

	waiter   sync.WaitGroup
	msgsChan []chan *KafkaData
}

type KafkaData struct {
	Uid int64 `json:"uid"`
}

func NewService(c config.Config) *Service {
	s := &Service{
		c:        c,
		UserRpc:  userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		msgsChan: make([]chan *KafkaData, chanCount),
	}
	for i := 0; i < chanCount; i++ {
		ch := make(chan *KafkaData, bufferCount)
		s.msgsChan[i] = ch
		s.waiter.Add(1)
		go s.consume(ch)
		//go s.consumeDTM(ch)
	}

	return s
}

func (s *Service) consume(ch chan *KafkaData) {
	defer s.waiter.Done()

	for {
		m, ok := <-ch
		if !ok {
			log.Fatal("seckill rmq exit")
		}
		fmt.Printf("consume msg: %+v\n", m)
		////检查今天用户是否已经领取
		//_, err := s.ProductRPC.CheckAndUpdateStock(context.Background(), &product.CheckAndUpdateStockRequest{ProductId: m.Pid})
		//if err != nil {
		//	logx.Errorf("s.ProductRPC.CheckAndUpdateStock pid: %d error: %v", m.Pid, err)
		//	return
		//}
		//_, err = s.OrderRPC.CreateOrder(context.Background(), &order.CreateOrderRequest{Uid: m.Uid, Pid: m.Pid})
		//if err != nil {
		//	logx.Errorf("CreateOrder uid: %d pid: %d error: %v", m.Uid, m.Pid, err)
		//}
	}
}

//var dtmServer = "etcd://127.0.0.1:2379/dtmservice"
//
//// 消费 Kafka 消息并执行分布式事务处理
//func (s *Service) consumeDTM(ch chan *KafkaData) {
//	defer s.waiter.Done()
//
//	//productServer, err := s.c.ProductRPC.BuildTarget()
//	//if err != nil {
//	//	log.Fatalf("s.c.ProductRPC.BuildTarget error: %v", err)
//	//}
//	//orderServer, err := s.c.OrderRPC.BuildTarget()
//	//if err != nil {
//	//	log.Fatalf("s.c.OrderRPC.BuildTarget error: %v", err)
//	//}
//
//	for {
//		m, ok := <-ch
//		if !ok {
//			log.Fatal("seckill rmq exit")
//		}
//		fmt.Printf("consume msg: %+v\n", m)
//		fmt.Printf("Connecting to: %s\n", dtmServer)
//
//		// 生成全局事务 ID
//		gid := dtmgrpc.MustGenGid(dtmServer)
//		err := dtmgrpc.TccGlobalTransaction(dtmServer, gid, func(tcc *dtmgrpc.TccGrpc) error {
//			fmt.Println("321123")
//			//// 调用商品服务进行事务处理
//			//if e := tcc.CallBranch(); err != nil {
//			//	logx.Errorf("tcc.CallBranch server: %s error: %v", productServer, err)
//			//	return e
//			//}
//			//// 调用订单服务进行事务处理
//			//if e := tcc.CallBranch(
//			//	&order.CreateOrderRequest{Uid: m.Uid, Pid: m.Pid},
//			//	orderServer+"/order.Order/CreateOrderCheck",
//			//	orderServer+"/order.Order/CreateOrder",
//			//	orderServer+"/order.Order/RollbackOrder",
//			//	&order.CreateOrderResponse{},
//			//); err != nil {
//			//	logx.Errorf("tcc.CallBranch server: %s error: %v", orderServer, err)
//			//	return e
//			//}
//			return nil
//		})
//		logger.FatalIfError(err)
//	}
//}

// Consume 函数用于处理接收到的消息，并将其发送到对应的消息通道中
func (s *Service) Consume(_ string, value string) error {
	logx.Infof("Consume value: %s\n", value)
	var data []*KafkaData
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return err
	}
	for _, d := range data {
		index := d.Uid % int64(chanCount)
		fmt.Printf("consume msg index: %d\n, d.Uid: %d\n", index, d.Uid)
		s.msgsChan[index] <- d

	}
	return nil
}
