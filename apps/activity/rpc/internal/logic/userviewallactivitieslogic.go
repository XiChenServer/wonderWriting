package logic

import (
	"calligraphy/apps/activity/model"
	"context"
	"encoding/json"
	"fmt"

	"calligraphy/apps/activity/rpc/internal/svc"
	"calligraphy/apps/activity/rpc/types/activity"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserViewAllActivitiesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserViewAllActivitiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserViewAllActivitiesLogic {
	return &UserViewAllActivitiesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserViewAllActivities 用户查看自己的报名活动
func (l *UserViewAllActivitiesLogic) UserViewAllActivities(in *activity.UserViewAllActivitiesRequest) (*activity.UserViewAllActivitiesResponse, error) {
	// 获取请求中的用户ID和分页参数
	userID := in.UserId
	pageIndex := int(in.Page)
	pageSize := int(in.PageSize)
	start := (pageIndex - 1) * pageSize

	// 构建缓存键
	cacheKey := fmt.Sprintf("user_activities:%d:%d:%d", userID, pageIndex, pageSize)

	// 尝试从缓存中获取数据
	cachedData, err := getFromCacheUserActivities(l, cacheKey)
	if err == nil {
		return cachedData, nil // 缓存命中，直接返回缓存数据
	}

	// 缓存未命中，查询数据库获取数据
	userActivities, err := (&model.UserSignUpActivity{}).GetUserActivities(l.svcCtx.DB, uint(userID), start, pageSize)
	if err != nil {
		return nil, err
	}

	// 查询用户参加的活动信息
	activitiesInfo := make([]*activity.ActivityInfo, 0)
	for _, activityID := range userActivities {
		activityInfo, err := (&model.Activity{}).GetActivityInfo(l.svcCtx.DB, activityID)
		if err != nil {
			return nil, err
		}
		newInfo := &activity.ActivityInfo{
			Id:           uint32(activityInfo.ID),
			Name:         activityInfo.Name,
			ActivityInfo: activityInfo.ActivityInfo,
			Location:     activityInfo.Location,
			DateTime:     activityInfo.DateTime,
			Organizer:    activityInfo.Organizer,
			EndDateTime:  activityInfo.EndDateTime,
			Duration:     activityInfo.Duration,
			RewardsInfo:  activityInfo.RewardsInfo,
		}
		activitiesInfo = append(activitiesInfo, newInfo)
	}

	// 查询总活动数和总页数
	totalActivities, err := (&model.UserSignUpActivity{}).GetUserActivitiesCount(l.svcCtx.DB, uint(userID))
	if err != nil {
		return nil, err
	}
	totalPages := totalActivities / pageSize
	if totalActivities%pageSize != 0 {
		totalPages++
	}

	// 构建并返回活动信息响应，包括分页相关信息
	resp := &activity.UserViewAllActivitiesResponse{
		Activities:  activitiesInfo,
		CurrentPage: uint32(pageIndex),
		PageSize:    uint32(pageSize),
		Offset:      uint32(start),
		TotalCount:  uint64(totalActivities), // 添加总活动数字段
		TotalPages:  uint32(totalPages),      // 添加总页数字段
		Overflow:    pageIndex > totalPages,
	}

	// 将查询结果存入缓存
	cacheTime := 60 * 5 // 缓存过期时间为 5 分钟
	err = setCacheUserActivities(l, cacheKey, resp, cacheTime)
	if err != nil {
		fmt.Println("Failed to set cache:", err)
	}

	return resp, nil
}

// 从缓存中获取用户活动数据
func getFromCacheUserActivities(l *UserViewAllActivitiesLogic, key string) (*activity.UserViewAllActivitiesResponse, error) {
	val, err := l.svcCtx.RDB.GetCtx(l.ctx, key)
	if err != nil {
		return nil, err
	}

	var data activity.UserViewAllActivitiesResponse
	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// 将用户活动数据存入缓存
func setCacheUserActivities(l *UserViewAllActivitiesLogic, key string, data *activity.UserViewAllActivitiesResponse, ttl int) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = l.svcCtx.RDB.SetexCtx(l.ctx, key, string(jsonData), ttl)
	if err != nil {
		return err
	}

	return nil
}
