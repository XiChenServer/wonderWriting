package logic

import (
	postModel "calligraphy/apps/community/model"
	"calligraphy/apps/home/rpc/types/home"
	userModel "calligraphy/apps/user/model"
	"context"
	"encoding/json"
	"log"
	"time"

	"calligraphy/apps/home/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type PostPopularityRankingsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostPopularityRankingsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostPopularityRankingsLogic {
	return &PostPopularityRankingsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

var (
	lastUpdated time.Time
	cacheKey    = "post_popularity_rankings"
)

// UpdateCachePost 更新排行榜缓存数据
func (l *PostPopularityRankingsLogic) UpdateCachePost() {
	// 从model层获取数据
	res, err := (&postModel.Post{}).GetTopLikedPosts(l.svcCtx.DB)
	if err != nil {
		log.Println("Failed to get data from database:", err)
		return
	}

	// 对数据进行转换
	var postPopularData []*home.PostPopularityInfo
	for _, v := range res {
		newPostPopularData := &home.PostPopularityInfo{
			PostId:          uint32(v.ID),
			Content:         v.Content,
			LikeCont:        int64(v.LikeCount),
			CollectionCount: int64(v.CollectionCount),
			CommentCount:    int64(v.CommentCount),
		}

		// 获取用户信息
		userRes, err := (&userModel.User{}).FindOne(l.svcCtx.DB, v.UserID)
		if err != nil {
			log.Println("Failed to get user data from database:", err)
		} else {
			newUserInfo := &home.UserPopularInfo{
				UserId:    uint32(userRes.UserID),
				NickName:  userRes.Nickname,
				Account:   userRes.Account,
				LikeCount: int64(userRes.LikeCount),
				Avatar:    "qiniu.ImgUrl" + userRes.AvatarBackground,
			}
			newPostPopularData.PopularInfo = newUserInfo
		}

		// 获取帖子图片信息
		postImage, err := (&postModel.PostImage{}).FindImageByPostId(l.svcCtx.DB, v.ID)
		if err != nil {
			log.Println("Failed to get post image data from database:", err)
		} else {
			newPostPopularData.PostImage = postImage
		}

		postPopularData = append(postPopularData, newPostPopularData)
	}

	// 构建响应
	response := &home.PostPopularityRankingsResponse{PostPopularData: postPopularData}

	// 将响应数据存入缓存中
	responseJSON, err := json.Marshal(response)
	if err != nil {
		log.Println("Failed to marshal response data:", err)
		return
	}

	// 更新数据的更新时间
	lastUpdated = time.Now()

	// 缓存时间为5分钟
	cacheTime := 5 * time.Minute
	err = l.svcCtx.RDB.SetexCtx(l.ctx, cacheKey, string(responseJSON), int(cacheTime.Seconds()))
	if err != nil {
		log.Println("Failed to set cache data:", err)
		return
	}
}

// PostPopularityRankings 获取排行榜数据（带分页）
func (l *PostPopularityRankingsLogic) PostPopularityRankings(in *home.PostPopularityRankingsRequest) (*home.PostPopularityRankingsResponse, error) {
	// 检查是否需要更新缓存
	if time.Since(lastUpdated) > 10*time.Minute {
		l.UpdateCachePost()
	}

	// 获取缓存数据
	value, err := l.svcCtx.RDB.GetCtx(l.ctx, cacheKey)
	if err != nil {
		log.Println("Failed to get cache data:", err)
		return nil, err
	}

	// 解析缓存数据
	var response home.PostPopularityRankingsResponse
	err = json.Unmarshal([]byte(value), &response)
	if err != nil {
		log.Println("Failed to unmarshal cached data:", err)
		return nil, err
	}

	// 根据请求参数进行分页
	pageIndex := int(in.Page)
	pageSize := int(in.PageSize)
	totalPosts := len(response.PostPopularData)

	// 计算总页数
	totalPages := totalPosts / pageSize
	if totalPosts%pageSize != 0 {
		totalPages++
	}

	// 计算当前页的起始和结束位置
	start := (pageIndex - 1) * pageSize
	end := start + pageSize

	if start >= totalPosts {
		// 请求的起始位置超出了数据范围，返回空数据
		return &home.PostPopularityRankingsResponse{
			PostPopularData: []*home.PostPopularityInfo{},
			PageSize:        uint32(pageSize),
			CurrentPage:     uint32(pageIndex),
			TotalPages:      uint32(totalPages),
			TotalCount:      uint64(totalPosts),

			Offset:   uint32(start),
			Overflow: true,
		}, nil
	}

	if end > totalPosts {
		end = totalPosts
	}

	// 截取分页数据
	pagedData := response.PostPopularData[start:end]

	return &home.PostPopularityRankingsResponse{
		PageSize:        uint32(pageSize),
		CurrentPage:     uint32(pageIndex),
		TotalPages:      uint32(totalPages),
		TotalCount:      uint64(totalPosts),
		PostPopularData: pagedData,
		Offset:          uint32(start),
		Overflow:        false,
	}, nil
}
