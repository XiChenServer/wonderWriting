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

var totalUsers int // 全局变量，用于保存用户总数

// UpdateCacheUser 更新用户人气排行榜缓存
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
	response := &home.UserPopularityRankingsResponse{
		UserPopularData: userPopularData,
		TotalCount:      uint64(len(userPopularData)), // 计算用户总数
	}

	// 将响应数据存入缓存中
	responseData, err := json.Marshal(response)
	if err != nil {
		log.Println("Failed to marshal response data:", err)
		return
	}

	// 更新数据的更新时间
	userlastUpdated = time.Now()

	// 缓存时间为1小时
	cacheTime := 1 * time.Hour
	err = l.svcCtx.RDB.SetexCtx(l.ctx, "user_popularity_rankings", string(responseData), int(cacheTime.Seconds()))
	if err != nil {
		log.Println("Failed to set cache data:", err)
		return
	}

	// 更新全局用户总数变量
	totalUsers = len(userPopularData)
}

// UserPopularityRankings 获取用户人气排行榜
func (l *UserPopularityRankingsLogic) UserPopularityRankings(in *home.UserPopularityRankingsRequest) (*home.UserPopularityRankingsResponse, error) {
	fmt.Println("123")
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

	// 设置总用户数
	response.TotalCount = uint64(totalUsers)

	// 根据请求参数进行分页
	pageIndex := int(in.Page)
	pageSize := int(in.PageSize)
	totalUsers := len(response.UserPopularData)

	// 计算总页数
	totalPages := totalUsers / pageSize
	if totalUsers%pageSize != 0 {
		totalPages++
	}

	// 计算当前页的起始和结束位置
	start := (pageIndex - 1) * pageSize
	end := start + pageSize

	if start >= totalUsers {
		// 请求的起始位置超出了数据范围，返回空数据
		return &home.UserPopularityRankingsResponse{
			PageSize:        uint32(pageSize),
			CurrentPage:     uint32(pageIndex),
			TotalPages:      uint32(totalPages),
			TotalCount:      uint64(totalUsers),
			UserPopularData: []*home.UserPopularInfo{},
			Offset:          uint32(start),
			Overflow:        true,
		}, nil
	}

	if end > totalUsers {
		end = totalUsers
	}

	// 截取分页数据
	pagedData := response.UserPopularData[start:end]

	return &home.UserPopularityRankingsResponse{
		PageSize:        uint32(pageSize),
		CurrentPage:     uint32(pageIndex),
		TotalPages:      uint32(totalPages),
		TotalCount:      uint64(totalUsers),
		UserPopularData: pagedData,
		Offset:          uint32(start),
		Overflow:        false,
	}, nil
}
