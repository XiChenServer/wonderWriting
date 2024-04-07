package logic

import (
	postModel "calligraphy/apps/community/model"
	userModel "calligraphy/apps/user/model"
	"calligraphy/pkg/qiniu"
	"context"
	"encoding/json"
	"log"
	"time"

	"calligraphy/apps/home/rpc/internal/svc"
	"calligraphy/apps/home/rpc/types/home"

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

func UpdateCachePost(l *PostPopularityRankingsLogic) {
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
				Avatar:    qiniu.ImgUrl + userRes.AvatarBackground,
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
	responseData, err := json.Marshal(response)
	if err != nil {
		log.Println("Failed to marshal response data:", err)
		return
	}

	postlastUpdated = time.Now() // 更新数据的更新时间
	seconds := int(time.Minute.Seconds())
	err = l.svcCtx.RDB.SetexCtx(l.ctx, "post_popularity_rankings", string(responseData), seconds)
	if err != nil {
		log.Println("Failed to set cache data:", err)
		return
	}
}

func (l *PostPopularityRankingsLogic) PostPopularityRankings(in *home.PostPopularityRankingsRequest) (*home.PostPopularityRankingsResponse, error) {
	// 每分钟更新一次缓存数据
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	// 检查是否需要更新缓存
	if time.Since(postlastUpdated) > time.Minute {
		UpdateCachePost(l)
	}

	// 使用 Redis 客户端的 Get 方法获取键对应的值
	value, err := l.svcCtx.RDB.GetCtx(l.ctx, "post_popularity_rankings")
	if err != nil {
		log.Println("Failed to get cache data:", err)
		return nil, err
	}

	log.Println("Cached data:", value)

	// 检查获取到的数据是否为空
	if value == "" {
		log.Println("Cached data is empty")
		// 重新刷新缓存数据
		UpdateCachePost(l)

		// 再次获取最新的缓存数据
		value, err = l.svcCtx.RDB.GetCtx(l.ctx, "post_popularity_rankings")
		if err != nil {
			log.Println("Failed to get cache data after refresh:", err)
			return nil, err
		}
	}

	// 返回缓存数据
	var response home.PostPopularityRankingsResponse
	err = json.Unmarshal([]byte(value), &response)
	if err != nil {
		log.Println("Failed to unmarshal cached data:", err)
		return nil, err
	}

	return &response, nil
}
