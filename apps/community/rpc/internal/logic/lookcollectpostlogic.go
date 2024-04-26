package logic

import (
	"calligraphy/apps/community/model"
	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type LookCollectPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLookCollectPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookCollectPostLogic {
	return &LookCollectPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LookCollectPostLogic) LookCollectPost(in *community.LookCollectPostRequest) (*community.LookCollectPostResponse, error) {
	// todo: add your logic here and delete this line
	res, err := (&model.Collect{}).FindCollect(l.svcCtx.DB, uint(in.UserId))
	if err != nil {
		l.Error("rpc 用户查看自己收藏的帖子的时候出现问题，可能是没有数据，可能是数据库查询出现了问题，err:", err.Error())
		return nil, err
	}

	var posts []*community.PostInfo
	for _, v := range *res {
		// 查询每个帖子的图片信息
		var urls []string
		urls, err = (&model.PostImage{}).FindImageByPostId(l.svcCtx.DB, v.PostID)
		if err != nil {
			return nil, err
		}

		// 查询用户信息
		userInfo, err := getUserInfo(l.svcCtx.DB, int(v.UserID))

		if err != nil {
			return nil, err
		}
		postInfoData, err := (&model.Post{}).LookPostByPostId(l.svcCtx.DB, v.PostID)
		// 创建新的帖子信息结构体
		newPost := &community.PostInfo{
			Id:           uint32(v.ID),
			UserId:       uint32(v.UserID),
			LikeCount:    uint32(postInfoData.LikeCount),
			Content:      postInfoData.Content,
			CreateTime:   uint32(postInfoData.CreatedAt.Unix()),
			ImageUrls:    urls,
			CollectCount: uint32(postInfoData.CollectionCount),
			ContentCount: uint32(postInfoData.CommentCount),
			UserInfo:     userInfo,
		}

		// 将新的帖子信息添加到切片中
		posts = append(posts, newPost)
	}
	l.Info("rpc 用户查看自己的收藏帖子成功。userId", in.UserId)
	return &community.LookCollectPostResponse{
		PostData: posts,
	}, nil
}
