package logic

import (
	userModel "calligraphy/apps/user/model"
	"calligraphy/pkg/qiniu"
	"context"
	"encoding/json"
	"fmt"

	"calligraphy/apps/user/rpc/internal/svc"
	"calligraphy/apps/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LookAllFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLookAllFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookAllFollowLogic {
	return &LookAllFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户查看自己的关注（带缓存和分页）
func (l *LookAllFollowLogic) LookAllFollow(in *user.LookAllFollowRequest) (*user.LookAllFollowResponse, error) {
	// 构建缓存键
	cacheKey := fmt.Sprintf("user_follow:%d:%d:%d", in.UserId, in.Page, in.PageSize)

	// 尝试从缓存中获取数据
	cachedData, err := getFromCacheForFollow(l, cacheKey)
	if err == nil {
		return cachedData, nil // 缓存命中，直接返回缓存数据
	}

	// 缓存未命中，查询数据库获取数据
	res, err := (&userModel.Follow{}).LookAllFollow(l.svcCtx.DB, uint(in.UserId))
	if err != nil {
		return nil, err
	}

	var userInfo []*user.UserInfo
	for _, v := range *res {
		userinfo, err := (&userModel.User{}).FindOne(l.svcCtx.DB, v.FollowedUserID)
		if err != nil {
			return nil, err
		}
		newUserInfo := user.UserInfo{
			UserId:           uint32(userinfo.UserID),
			AvatarBackground: qiniu.ImgUrl + userinfo.AvatarBackground,
			NickName:         userinfo.Nickname,
			FollowCount:      int64(userinfo.FollowCount),
			FansCount:        int64(userinfo.FansCount),
			Email:            userinfo.Email,
		}
		userInfo = append(userInfo, &newUserInfo)
	}

	// 计算总页数
	totalCount := len(userInfo)
	totalPages := totalCount / int(in.PageSize)
	if totalCount%int(in.PageSize) != 0 {
		totalPages++
	}

	// 分页处理
	startIndex := (int(in.Page) - 1) * int(in.PageSize)
	endIndex := startIndex + int(in.PageSize)
	if endIndex > totalCount {
		endIndex = totalCount
	}
	pagedUserInfo := userInfo[startIndex:endIndex]

	// 构建并返回关注信息响应
	resp := &user.LookAllFollowResponse{
		UserInfo:    pagedUserInfo,
		CurrentPage: in.Page,
		PageSize:    in.PageSize,
		Offset:      uint32(startIndex),
		Overflow:    in.Page > uint32(totalPages),
		TotalPages:  uint32(totalPages),
		TotalCount:  uint64(totalCount),
	}

	cacheTime := 60 * 5 // 缓存时间为5分钟
	// 将查询结果存入缓存
	err = l.svcCtx.RDB.SetexCtx(l.ctx, cacheKey, toJson(resp), cacheTime)
	if err != nil {
		fmt.Println("Failed to set cache:", err)
	}

	return resp, nil
}

// 从缓存中获取关注数据
func getFromCacheForFollow(l *LookAllFollowLogic, key string) (*user.LookAllFollowResponse, error) {
	val, err := l.svcCtx.RDB.GetCtx(l.ctx, key)
	if err != nil {
		return nil, err
	}

	var data user.LookAllFollowResponse
	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
