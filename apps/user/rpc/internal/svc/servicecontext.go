package svc

import (
	"calligraphy/apps/user/model"
	"calligraphy/apps/user/rpc/internal/config"
	"fmt"
	"github.com/8treenet/gcache"
	"github.com/8treenet/gcache/option"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open("mysql", getDSN(&c))

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&model.User{}, &model.Follow{})

	opt := option.DefaultOption{}
	opt.Expires = 300              //缓存时间, 默认120秒。范围30-43200
	opt.Level = option.LevelSearch //缓存级别，默认LevelSearch。LevelDisable:关闭缓存，LevelModel:模型缓存， LevelSearch:查询缓存
	opt.AsyncWrite = false         //异步缓存更新, 默认false。 insert update delete 成功后是否异步更新缓存。 ps: affected如果未0，不触发更新。
	opt.PenetrationSafe = false    //开启防穿透, 默认false。 ps:防击穿强制全局开启。

	//缓存中间件附加到gorm.DB
	gcache.AttachDB(db, &opt, &option.RedisOption{Addr: "localhost:6379"})

	return &ServiceContext{
		Config: c,
		DB:     db,
		//UserModel: model.NewUsersModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
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
