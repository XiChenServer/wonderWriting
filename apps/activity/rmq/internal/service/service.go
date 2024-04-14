package service

import (
	"calligraphy/apps/activity/rmq/internal/config"
	"calligraphy/apps/activity/rpc/activityclient"
	userModel "calligraphy/apps/user/model"
	"calligraphy/apps/user/rpc/userclient"
	"calligraphy/common/app_redis"
	"encoding/json"
	"fmt"
	"github.com/8treenet/gcache"
	"github.com/8treenet/gcache/option"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"log"
	"sync"
	"time"
)

const (
	chanCount   = 10
	bufferCount = 1024
)

type Service struct {
	c           config.Config
	UserRpc     userclient.User
	DB          *gorm.DB
	RDB         *redis.Redis
	waiter      sync.WaitGroup
	msgsChan    []chan *KafkaData
	ActivityRpc activityclient.Activity
}

type KafkaData struct {
	Uid int64 `json:"uid"`
}

// 定义一个RPC接口，用于向用户发送消息
type MessageSender interface {
	SendMessage(userID uint64, message string) error
}

func NewService(c config.Config) *Service {
	//连接数据库
	db, err := gorm.Open("mysql", getDSN(&c))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	//初始化程序
	db.AutoMigrate()
	opt := option.DefaultOption{}
	opt.Expires = 300              //缓存时间, 默认120秒。范围30-43200
	opt.Level = option.LevelSearch //缓存级别，默认LevelSearch。LevelDisable:关闭缓存，LevelModel:模型缓存， LevelSearch:查询缓存
	opt.AsyncWrite = false         //异步缓存更新, 默认false。 insert update delete 成功后是否异步更新缓存。 ps: affected如果未0，不触发更新。
	opt.PenetrationSafe = false    //开启防穿透, 默认false。 ps:防击穿强制全局开启。

	//缓存中间件附加到gorm.DB
	gcache.AttachDB(db, &opt, &option.RedisOption{Addr: "localhost:6379"})

	s := &Service{
		c:           c,
		UserRpc:     userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		msgsChan:    make([]chan *KafkaData, chanCount),
		DB:          db,
		RDB:         app_redis.Redis,
		ActivityRpc: activityclient.NewActivity(zrpc.MustNewClient(c.ActivityRpc)),
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

// kafka进行抢积分抢积分的活动
func (s *Service) consume(ch chan *KafkaData) {
	defer s.waiter.Done()
	for {
		m, ok := <-ch
		if !ok {
			log.Fatal("seckill rmq exit")
		}
		fmt.Printf("consume msg: %+v\n", m)
		////检查今天用户是否已经领取
		// 检查用户是否已领取
		claimed, err := CheckUserClaimed(s.RDB, uint(m.Uid))
		if err != nil {
			logx.Errorf("failed to check user claimed: %v", err)
			return
		}
		if claimed {
			fmt.Println("123")
			// 用户已领取，跳过处理
			return
		}
		// 更新用户领取状态和积分数
		if err = (&userModel.User{}).UpdatePointsGrab(s.DB, uint(m.Uid)); err != nil {
			logx.Errorf("failed to update user points: %v", err)
			return
		}
		if err = SetUserClaimed(s.RDB, uint(m.Uid)); err != nil {
			logx.Errorf("set redis err: %v", err)
			return
		}
		if _, err = s.RDB.Decrby("point_count_one_day", 1); err != nil {
			logx.Errorf("Decrby redis err: %v", err)
			return
		}

	}
}

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

func getDSN(c *config.Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		c.MySQL.User,
		c.MySQL.Password,
		c.MySQL.Host,
		c.MySQL.Port,
		c.MySQL.Database,
	)
}

// CheckUserClaimed 检查其中有没有存在于redis
func CheckUserClaimed(redis *redis.Redis, userID uint) (bool, error) {
	key := fmt.Sprintf("user:%d:claimed", userID)
	exists, err := redis.Exists(key)
	fmt.Println("1")
	if err != nil {
		return false, err
	}
	if exists == true {
		return true, nil
	}
	return false, nil
}

// SetUserClaimed 将数据放到redis
func SetUserClaimed(redis *redis.Redis, userID uint) error {
	key := fmt.Sprintf("user:%d:claimed", userID)
	value := "true"

	expiration := 15 * time.Second // 过期时间为一天
	return redis.Setex(key, value, int(expiration.Seconds()))
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
