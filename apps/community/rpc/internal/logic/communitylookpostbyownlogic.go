package logic

import (
	"calligraphy/apps/community/model"
	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"

	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommunityLookPostByOwnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommunityLookPostByOwnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommunityLookPostByOwnLogic {
	return &CommunityLookPostByOwnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CommunityLookPostByOwn 查看某人的帖子（带缓存和分页）
func (l *CommunityLookPostByOwnLogic) CommunityLookPostByOwn(in *community.CommunityLookPostByOwnRequest) (*community.CommunityLookPostByOwnResponses, error) {
	// 构建缓存键
	cacheKey := fmt.Sprintf("user_posts:%d:%d:%d", in.UserId, in.Page, in.PageSize)

	// 尝试从缓存中获取数据
	cachedData, err := getFromCacheOneSelf(l, cacheKey)
	if err == nil {
		return cachedData, nil // 缓存命中，直接返回缓存数据
	}

	// 缓存未命中，查询数据库获取数据
	postInfo, totalPage, err := getPostsByUser(l.svcCtx.DB, in.UserId, int(in.Page), int(in.PageSize))
	if err != nil {
		return nil, err
	}

	// 构建并返回帖子信息响应
	resp := &community.CommunityLookPostByOwnResponses{
		PostData:    postInfo,
		CurrentPage: in.Page,
		PageSize:    in.PageSize,
		Offset:      (in.Page - 1) * in.PageSize,
		Overflow:    in.Page > totalPage,
		TotalPages:  totalPage,
		TotalCount:  uint64(len(postInfo)), // 假设这里的 TotalCount 表示当前页的帖子数量
	}

	cacheTime := 60 * 5 // 缓存时间为5分钟
	// 将查询结果存入缓存
	err = l.svcCtx.RDB.SetexCtx(l.ctx, cacheKey, toJson(resp), cacheTime)
	if err != nil {
		fmt.Println("Failed to set cache:", err)
	}

	return resp, nil
}

// 查询用户的帖子信息并进行分页
func getPostsByUser(db *gorm.DB, userID uint32, page, pageSize int) ([]*community.PostInfo, uint32, error) {
	var res []model.Post

	offset := (page - 1) * pageSize
	err := db.Where("user_id = ?", userID).Offset(offset).Limit(pageSize).Find(&res).Error
	if err != nil {
		return nil, 0, err
	}

	var totalCount int64
	err = db.Model(&model.Post{}).Where("user_id = ?", userID).Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}

	totalPage := uint32(totalCount) / uint32(pageSize)
	if uint32(totalCount)%uint32(pageSize) != 0 {
		totalPage++
	}

	// 构建用于返回的帖子信息切片
	var postInfo []*community.PostInfo

	// 遍历查询到的帖子信息
	for _, v := range res {
		// 查询每个帖子的图片信息
		var urls []string
		urls, err = (&model.PostImage{}).FindImageByPostId(db, v.ID)
		if err != nil {
			return nil, 0, err
		}

		// 查询用户信息
		userInfo, err := getUserInfo(db, int(v.UserID))
		if err != nil {
			return nil, 0, err
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

		// 将新的帖子信息添加到切片中
		postInfo = append(postInfo, newPost)
	}

	return postInfo, totalPage, nil
}

// 从缓存中获取数据
func getFromCacheOneSelf(l *CommunityLookPostByOwnLogic, key string) (*community.CommunityLookPostByOwnResponses, error) {
	val, err := l.svcCtx.RDB.GetCtx(l.ctx, key)
	if err != nil {
		return nil, err
	}

	var data community.CommunityLookPostByOwnResponses
	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
