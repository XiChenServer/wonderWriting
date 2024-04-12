package logic

import (
	activityModel "calligraphy/apps/activity/model"
	"calligraphy/apps/activity/rpc/internal/svc"
	"calligraphy/apps/activity/rpc/types/activity"
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type LookAllActivitiesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLookAllActivitiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookAllActivitiesLogic {
	return &LookAllActivitiesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func (l *LookAllActivitiesLogic) LookAllActivities(in *activity.LookAllActivitiesRequest) (*activity.LookAllActivitiesResponse, error) {
	// 提取分页参数
	pageIndex := int(in.Page)
	pageSize := int(in.PageSize)
	start := (pageIndex - 1) * pageSize

	// 构建缓存键
	cacheKey := fmt.Sprintf("activities:%d:%d", pageIndex, pageSize)

	// 尝试从缓存中获取数据
	cachedData, err := getFromCacheLookAllActivities(l, cacheKey)
	if err == nil {
		return cachedData, nil // 缓存命中，直接返回缓存数据
	}

	// 缓存未命中，查询数据库获取数据
	// 这里可以根据具体的业务逻辑从数据库中获取活动数据，并计算总活动数和总页数
	// 这里使用示例数据，实际情况下需要根据实际业务进行修改
	totalActivities := 1000 // 假设总活动数为 1000
	totalPages := totalActivities / pageSize
	if totalActivities%pageSize != 0 {
		totalPages++
	}

	// 构建用于返回的活动数据切片
	activities := make([]*activity.ActivityInfo, 0)

	// 查询数据库获取当前页的活动数据，假设从数据库中获取活动数据的函数为 GetAllActivities
	allActivities, err := (&activityModel.Activity{}).GetAllActivities(l.svcCtx.DB, start, pageSize)
	if err != nil {
		return nil, err
	}
	// 将查询到的活动数据转换为 activity.ActivityInfo 结构体，并添加到 activities 切片中
	for _, a := range allActivities {
		activityInfo := &activity.ActivityInfo{
			Id:           uint32(a.ID),
			Name:         a.Name,
			ActivityInfo: a.ActivityInfo,
			Location:     a.Location,
			DateTime:     a.DateTime,
			Organizer:    a.Organizer,
			EndDateTime:  a.EndDateTime,
			Duration:     a.Duration,
			RewardsInfo:  a.RewardsInfo,
		}
		activities = append(activities, activityInfo)
	}

	// 构建并返回活动信息响应，包括分页相关信息
	resp := &activity.LookAllActivitiesResponse{
		Activities:  activities,
		PageSize:    uint32(pageSize),
		CurrentPage: uint32(pageIndex),
		TotalPages:  uint32(totalPages),
		TotalCount:  uint64(totalActivities),
		Offset:      uint32(start),
		Overflow:    pageIndex > totalPages,
	}

	// 将查询结果存入缓存
	cacheTime := 60 * 5 // 缓存过期时间为 5 分钟
	err = l.svcCtx.RDB.SetexCtx(l.ctx, cacheKey, toJson(resp), cacheTime)
	if err != nil {
		fmt.Println("Failed to set cache:", err)
	}

	return resp, nil
}

// 从缓存中获取活动数据
func getFromCacheLookAllActivities(l *LookAllActivitiesLogic, key string) (*activity.LookAllActivitiesResponse, error) {
	val, err := l.svcCtx.RDB.GetCtx(l.ctx, key)
	if err != nil {
		return nil, err
	}

	var data activity.LookAllActivitiesResponse
	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
func toJson(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}
