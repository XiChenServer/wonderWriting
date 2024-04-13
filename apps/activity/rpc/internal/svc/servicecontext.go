package svc

import (
	"calligraphy/apps/activity/model"
	"calligraphy/apps/activity/rpc/internal/config"
	"calligraphy/apps/user/rpc/userclient"
	"calligraphy/common/app_redis"
	"fmt"
	"github.com/8treenet/gcache"
	"github.com/8treenet/gcache/option"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"log"
)

type ServiceContext struct {
	Config      config.Config
	DB          *gorm.DB
	UserRpc     userclient.User
	RDB         *redis.Redis
	KafkaPusher *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open("mysql", getDSN(&c))

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db.AutoMigrate(&model.UserSignUpActivity{}, &model.Activity{})
	opt := option.DefaultOption{}
	opt.Expires = 300              //缓存时间, 默认120秒。范围30-43200
	opt.Level = option.LevelSearch //缓存级别，默认LevelSearch。LevelDisable:关闭缓存，LevelModel:模型缓存， LevelSearch:查询缓存
	opt.AsyncWrite = false         //异步缓存更新, 默认false。 insert update delete 成功后是否异步更新缓存。 ps: affected如果未0，不触发更新。
	opt.PenetrationSafe = false    //开启防穿透, 默认false。 ps:防击穿强制全局开启。
	//缓存中间件附加到gorm.DB
	gcache.AttachDB(db, &opt, &option.RedisOption{Addr: "localhost:6379"})

	return &ServiceContext{
		RDB:         app_redis.Redis,
		UserRpc:     userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		DB:          db,
		KafkaPusher: kq.NewPusher(c.Kafka.Addrs, c.Kafka.GraBPointsTopic),
		Config:      c,
	}
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
