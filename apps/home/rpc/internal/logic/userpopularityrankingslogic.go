package logic

import (
	"calligraphy/apps/home/rpc/internal/svc"
	"calligraphy/apps/home/rpc/types/home"
	userModel "calligraphy/apps/user/model"
	"calligraphy/pkg/qiniu"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPopularityRankingsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserPopularityRankingsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPopularityRankingsLogic {
	return &UserPopularityRankingsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

var (
	userlastUpdated time.Time
	postlastUpdated time.Time
	cachedData      []byte
)

func UpdateCacheUser(l *UserPopularityRankingsLogic) {
	// 从model层获取数据
	res, err := (&userModel.User{}).GetTopLikedUsers(l.svcCtx.DB)
	if err != nil {
		log.Println("Failed to get data from database:", err)
		return
	}

	// 对数据进行转换
	var userPopularData []*home.UserPopularInfo
	for _, v := range res {
		newUserPopularData := &home.UserPopularInfo{
			UserId:    uint32(v.UserID),
			NickName:  v.Nickname,
			Account:   v.Account,
			LikeCount: int64(v.LikeCount),
			Avatar:    qiniu.ImgUrl + v.AvatarBackground,
		}
		userPopularData = append(userPopularData, newUserPopularData)
	}

	// 构建响应
	response := &home.UserPopularityRankingsResponse{UserPopularData: userPopularData}

	// 将响应数据存入缓存中
	responseData, err := json.Marshal(response)
	if err != nil {
		log.Println("Failed to marshal response data:", err)
		return
	}

	cachedData = responseData
	userlastUpdated = time.Now() // 更新数据的更新时间
	seconds := int(time.Minute.Seconds())
	err = l.svcCtx.RDB.SetexCtx(l.ctx, "user_popularity_rankings", string(responseData), seconds)
	if err != nil {
		log.Println("Failed to set cache data:", err)
		return
	}
}

func (l *UserPopularityRankingsLogic) UserPopularityRankings(in *home.UserPopularityRankingsRequest) (*home.UserPopularityRankingsResponse, error) {
	// 每分钟更新一次缓存数据
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	// 检查是否需要更新缓存
	if time.Since(userlastUpdated) > time.Minute {
		UpdateCacheUser(l)
	}

	// 使用 Redis 客户端的 Get 方法获取键对应的值
	value, err := l.svcCtx.RDB.GetCtx(l.ctx, "user_popularity_rankings")
	if err != nil {
		log.Println("Failed to get cache data:", err)
		return nil, err
	}

	log.Println("Cached data:", value)

	// 检查获取到的数据是否为空
	if value == "" {
		log.Println("Cached data is empty")
		// 重新刷新缓存数据
		UpdateCacheUser(l)
		fmt.Println("1")
		// 再次获取最新的缓存数据
		value, err = l.svcCtx.RDB.GetCtx(l.ctx, "user_popularity_rankings")
		if err != nil {
			log.Println("Failed to get cache data after refresh:", err)
			return nil, err
		}
	}

	// 返回缓存数据
	var response home.UserPopularityRankingsResponse
	err = json.Unmarshal([]byte(value), &response)
	if err != nil {
		log.Println("Failed to unmarshal cached data:", err)
		return nil, err
	}

	return &response, nil
}
