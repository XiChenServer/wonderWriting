package logic

import (
	"calligraphy/apps/community/model"
	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"
	userModel "calligraphy/apps/user/model"
	"calligraphy/pkg/qiniu"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/zeromicro/go-zero/core/logx"
)

type CommunityLookAllPostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommunityLookAllPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommunityLookAllPostsLogic {
	return &CommunityLookAllPostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RPC方法：查看所有帖子（带缓存）
func (l *CommunityLookAllPostsLogic) CommunityLookAllPosts(in *community.CommunityLookAllPostsRequest) (*community.CommunityLookAllPostsResponse, error) {
	// 提取分页参数
	page := int(in.Page)
	pageSize := int(in.PageSize)

	// 构建缓存键
	cacheKey := fmt.Sprintf("posts:%d:%d", page, pageSize)

	// 尝试从缓存中获取数据
	cachedData, err := getFromCacheLookAll(l, cacheKey)
	if err == nil {

		return cachedData, nil // 缓存命中，直接返回缓存数据
	}

	// 缓存未命中，查询数据库获取数据
	posts, totalCount, err := (&model.Post{}).LookAllPostsWithPagination(l.svcCtx.DB, page, pageSize)
	for _, post := range posts {
		fmt.Println(post.Content)
	}
	if err != nil {
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
			return nil, err
		}

		// 查询用户信息
		userInfo, err := getUserInfo(l.svcCtx.DB, int(v.UserID))

		if err != nil {
			fmt.Println(v.UserID)
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
	resp := &community.CommunityLookAllPostsResponse{
		PostData:    postInfo,
		CurrentPage: uint32(page),
		PageSize:    uint32(pageSize),
		Offset:      uint32((page - 1) * pageSize),
		Overflow:    page > int(totalPages),
		TotalPages:  uint32(totalPages),
		TotalCount:  uint64(totalCount),
	}

	cacheTime := 60 * 5
	// 将查询结果存入缓存
	err = l.svcCtx.RDB.SetexCtx(l.ctx, cacheKey, toJson(resp), cacheTime) // 设置缓存过期时间为5分钟
	if err != nil {

		fmt.Println("Failed to set cache:", err)
	}

	return resp, nil
}

// 从缓存中获取数据
func getFromCacheLookAll(l *CommunityLookAllPostsLogic, key string) (*community.CommunityLookAllPostsResponse, error) {
	val, err := l.svcCtx.RDB.GetCtx(l.ctx, key)
	if err != nil {
		return nil, err
	}

	var data community.CommunityLookAllPostsResponse
	err = json.Unmarshal([]byte(val), &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// 查询用户信息
func getUserInfo(db *gorm.DB, userID int) (*community.UserSimpleInfo, error) {
	var user userModel.User
	fmt.Println(userID)
	err := db.Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	userInfo := &community.UserSimpleInfo{
		Id:          uint32(user.UserID),
		NickName:    user.Nickname,
		Account:     user.Account,
		AvatarImage: qiniu.ImgUrl + user.AvatarBackground,
	}

	return userInfo, nil
}

func toJson(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}
