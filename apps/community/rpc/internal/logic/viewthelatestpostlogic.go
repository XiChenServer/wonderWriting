package logic

import (
	"calligraphy/apps/community/model"
	"context"
	"encoding/json"
	"fmt"

	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewTheLatestPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewViewTheLatestPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewTheLatestPostLogic {
	return &ViewTheLatestPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ViewTheLatestPost 查询最新的帖子
func (l *ViewTheLatestPostLogic) ViewTheLatestPost(in *community.ViewTheLatestPostRequest) (*community.ViewTheLatestPostResponse, error) {
	// 提取分页参数
	page := int(in.Page)
	pageSize := int(in.PageSize)

	// 构建缓存键
	cacheKey := fmt.Sprintf("latest_posts:%d:%d", page, pageSize)

	// 尝试从缓存中获取数据
	cachedData, err := getFromCacheLatest(l, cacheKey)
	if err == nil {
		l.Logger.Infof("rpc ViewTheLatestPostLogic 从缓存中获取了数据 get cache key: %s", cacheKey)
		return cachedData, nil // 缓存命中，直接返回缓存数据
	}

	// 缓存未命中，查询数据库获取数据
	posts, totalCount, err := (&model.Post{}).LatestPostsWithPagination(l.svcCtx.DB, page, pageSize)
	if err != nil {
		l.Logger.Error("rpc LatestPostsWithPagination 缓存之中没有该页的缓存，在从数据库中获取信息的时候出现问题")
		return nil, err
	}

	// 计算总页数
	totalPages := totalCount / int64(pageSize)
	if totalCount%int64(pageSize) != 0 {
		totalPages++
	}

	// 构建用于返回的帖子信息切片
	postInfo := make([]*community.PostInfo, 0, len(posts))

	// 遍历查询到的帖子信息
	for _, v := range posts {
		// 查询每个帖子的图片信息
		urls, err := (&model.PostImage{}).FindImageByPostId(l.svcCtx.DB, v.ID)
		if err != nil {
			l.Logger.Error("rpc 在查找帖子的图片的时候，数据库操作出现了问题", err.Error())
			return nil, err
		}

		// 查询用户信息
		userInfo, err := getUserInfo(l.svcCtx.DB, int(v.UserID))
		if err != nil {
			l.Logger.Error("rpc 在查找帖子的用户的信息的时候，数据库操作出现了问题", err.Error())
			return nil, err
		}

		// 创建新的帖子信息结构体
		newPost := &community.PostInfo{
			Id:           uint32(v.ID),
			UserId:       uint32(v.UserID),
			LikeCount:    uint32(v.LikeCount),
			Content:      v.Content,
			CreateTime:   uint32(v.CreatedAt.Unix()),
			ImageUrls:    urls,
			CollectCount: uint32(v.CollectionCount),
			ContentCount: uint32(v.CommentCount),
			UserInfo:     userInfo,
		}

		//将新的帖子信息添加到切片中
		postInfo = append(postInfo, newPost)
	}

	// 构建并返回帖子信息响应，包括分页相关信息
	resp := &community.ViewTheLatestPostResponse{
		PostData:    postInfo,
		CurrentPage: uint32(page),
		PageSize:    uint32(pageSize),
		Offset:      uint32((page - 1) * pageSize),
		Overflow:    page > int(totalPages),
		TotalPages:  uint32(totalPages),
		TotalCount:  uint64(totalCount),
	}

	cacheTime := 30 // 5分钟
	// 将查询结果存入缓存
	err = l.svcCtx.RDB.SetexCtx(l.ctx, cacheKey, toJson(resp), cacheTime)
	if err != nil {
		l.Logger.Error("rpc 这次没有将拿到的信息存放到缓存里面去，", err.Error())
		fmt.Println("Failed to set cache:", err)
	}
	l.Logger.Info("rpc ViewTheLatestPost 成功获取到了信息，返回给上层")
	return resp, nil
}

// 从缓存中获取数据（最新帖子）
func getFromCacheLatest(l *ViewTheLatestPostLogic, key string) (*community.ViewTheLatestPostResponse, error) {
	val, err := l.svcCtx.RDB.GetCtx(l.ctx, key)
	if err != nil {
		return nil, err
	}

	var data community.ViewTheLatestPostResponse
	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		l.Logger.Error("在数据类型进行转换的时候出现了问题，err", err.Error())
		return nil, err
	}

	return &data, nil
}
