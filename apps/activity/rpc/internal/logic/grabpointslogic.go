package logic

import (
	"calligraphy/apps/activity/rpc/internal/svc"
	"calligraphy/apps/activity/rpc/types/activity"
	"calligraphy/pkg/batcher"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	limitPeriod       = 10
	limitQuota        = 1
	seckillUserPrefix = "seckill#u#"
	localCacheExpire  = time.Second * 60

	batcherSize     = 100
	batcherBuffer   = 100
	batcherWorker   = 10
	batcherInterval = time.Second
)

type GrabPointsLogic struct {
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	limiter    *limit.PeriodLimit
	localCache *collection.Cache
	batcher    *batcher.Batcher
	logx.Logger
}

func NewGrabPointsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrabPointsLogic {
	// 创建一个本地缓存，缓存的有效期为 localCacheExpire
	localCache, err := collection.NewCache(localCacheExpire)
	// 如果创建缓存时出现错误，则触发 panic，并打印错误信息
	if err != nil {
		panic(err)
	}
	// 创建 SeckillOrderLogic 结构体实例 s
	s := &GrabPointsLogic{
		// 传入上下文对象 ctx
		ctx: ctx,
		// 传入服务上下文对象 svcCtx
		svcCtx: svcCtx,
		// 使用带有上下文的日志对象
		Logger: logx.WithContext(ctx),
		// 设置本地缓存对象
		localCache: localCache,
		// 创建限流器对象，并设置相关参数
		limiter: limit.NewPeriodLimit(limitPeriod, limitQuota, svcCtx.RDB, seckillUserPrefix),
	}

	// 创建批处理器实例 b，并设置配置选项
	b := batcher.New(
		batcher.WithSize(batcherSize),         // 设置批处理器大小
		batcher.WithBuffer(batcherBuffer),     // 设置批处理器缓冲区大小
		batcher.WithWorker(batcherWorker),     // 设置批处理器工作线程数
		batcher.WithInterval(batcherInterval), // 设置批处理器间隔时间
	)

	// 设置批处理器的分片函数，用于计算分片
	b.Sharding = func(key string) int {
		// 将字符串键解析为整数并取模，得到分片结果
		pid, _ := strconv.ParseInt(key, 10, 64)
		return int(pid) % batcherWorker
	}

	// 设置批处理器的处理函数，用于处理批量消息
	b.Do = func(ctx context.Context, val map[string][]interface{}) {
		// 将批量消息转换为 JSON 格式
		var msgs []*KafkaData
		for _, vs := range val {
			for _, v := range vs {
				msgs = append(msgs, v.(*KafkaData))
			}
		}
		kd, err := json.Marshal(msgs)
		if err != nil {
			// 如果转换出错，则记录错误日志
			logx.Errorf("Batcher.Do json.Marshal msgs: %v error: %v", msgs, err)
		}
		// 将 JSON 格式的消息推送到 Kafka 中
		if err = s.svcCtx.KafkaPusher.Push(string(kd)); err != nil {
			// 如果推送出错，则记录错误日志
			logx.Errorf("KafkaPusher.Push kd: %s error: %v", string(kd), err)
		}
	}

	// 将批处理器 b 设置给 s.batcher，并启动批处理器
	s.batcher = b
	s.batcher.Start()

	// 返回初始化后的秒杀订单逻辑处理器实例 s
	return s

}

type KafkaData struct {
	Uid int64 `json:"uid"`
}

func (l *GrabPointsLogic) GrabPoints(in *activity.GrabPointsRequest) (*activity.GrabPointsResponse, error) {
	// 通过限流器检查是否可以继续抢积分
	code, _ := l.limiter.Take(strconv.FormatInt(int64(in.UserId), 10))
	if code == limit.OverQuota {
		return nil, status.Errorf(codes.OutOfRange, "Number of requests exceeded the limit")
	}
	// 从 Redis 中获取当日剩余积分数量
	count, err := l.svcCtx.RDB.Get("point_count_one_day")
	if err != nil {
		logx.Errorf("failed to get points count from Redis: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to get points count")
	}
	// 检查用户是否已领取
	claimed, err := CheckUserClaimed(l.svcCtx.RDB, uint(in.UserId))
	if err != nil {
		logx.Errorf("failed to check user claimed: %v", err)
		return &activity.GrabPointsResponse{}, err
	}
	if claimed == true {
		fmt.Println("123")
		// 用户已领取，跳过处理
		return &activity.GrabPointsResponse{}, errors.New("用户已经抢过了")
	}
	fmt.Println(in.UserId)
	countNum, _ := strconv.Atoi(count)
	// 检查当日剩余积分是否充足
	if countNum <= 0 {
		return nil, status.Errorf(codes.OutOfRange, "Insufficient stock")
	}

	// 向消息队列中添加抢积分请求
	if err := l.batcher.Add(strconv.FormatInt(int64(in.UserId), 10), &KafkaData{Uid: int64(in.UserId)}); err != nil {
		logx.Errorf("l.batcher.Add uid: %d error: %v", in.UserId, err)
	}
	time.Sleep(5 * time.Microsecond)
	// 检查用户是否已领取
	claimed, err = CheckUserClaimed(l.svcCtx.RDB, uint(in.UserId))
	if err != nil {
		logx.Errorf("failed to check user claimed: %v", err)
		return &activity.GrabPointsResponse{}, err
	}
	if claimed == false {
		fmt.Println("123")
		// 用户已领取，跳过处理
		return &activity.GrabPointsResponse{}, errors.New("没有抢到")
	}
	return &activity.GrabPointsResponse{}, nil
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
