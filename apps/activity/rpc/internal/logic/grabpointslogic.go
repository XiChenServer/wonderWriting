package logic

import (
	"calligraphy/apps/activity/rpc/internal/svc"
	"calligraphy/apps/activity/rpc/types/activity"
	"calligraphy/pkg/batcher"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/limit"
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
	ctx    context.Context
	svcCtx *svc.ServiceContext

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

var PointByOneDay = 10000

func (l *GrabPointsLogic) GrabPoints(in *activity.GrabPointsRequest) (*activity.GrabPointsResponse, error) {
	// 通过限流器检查是否可以继续抢积分
	code, _ := l.limiter.Take(strconv.FormatInt(int64(in.UserId), 10))
	if code == limit.OverQuota {
		return nil, status.Errorf(codes.OutOfRange, "Number of requests exceeded the limit")
	}

	// 检查当日剩余积分是否充足
	if PointByOneDay <= 0 {
		return nil, status.Errorf(codes.OutOfRange, "Insufficient stock")
	}

	// 向消息队列中添加抢积分请求
	if err := l.batcher.Add(strconv.FormatInt(int64(in.UserId), 10), &KafkaData{Uid: int64(in.UserId)}); err != nil {
		logx.Errorf("l.batcher.Add uid: %d error: %v", in.UserId, err)
	}

	return &activity.GrabPointsResponse{}, nil
}

func (l *GrabPointsLogic) UpdatePointsOneDay() bool {
	// 获取当前时间
	now := time.Now()

	// 确定每天更新积分的时间点（例如每天的凌晨零点）
	updateTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	// 判断当前时间是否已经过了更新时间，如果已经过了则表示今天已经更新过积分，无需再次更新
	if now.After(updateTime) {
		return false
	}

	// 获取剩余积分数量并更新 PointByOneDay 的值
	PointByOneDay = 10000

	// 返回更新成功
	return true
}
