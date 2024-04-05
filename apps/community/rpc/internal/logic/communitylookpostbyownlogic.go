package logic

import (
	"calligraphy/apps/community/model"
	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"
	"context"

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

func (l *CommunityLookPostByOwnLogic) CommunityLookPostByOwn(in *community.CommunityLookPostByOwnRequest) (*community.CommunityLookPostByOwnResponses, error) {
	// todo: add your logic here and delete this line
	// 创建 Post 和 PostImage 操作实例
	//postOperations := model.Post{}
	postImageOperations := model.PostImage{}
	var res []model.Post
	// 查询所有帖子信息
	//res, err := postOperations.LookPostByOwn(l.svcCtx.DB, uint(in.UserId))
	//if err != nil {
	//	return nil, err
	//}
	err := l.svcCtx.DB.Where("user_id = ?", in.UserId).Find(&res).Error
	if err != nil {
		return nil, err
	}
	// 创建用于返回的帖子信息切片
	var postInfo []*community.PostInfo

	// 遍历查询到的帖子信息
	for _, v := range res {
		// 查询每个帖子的图片信息
		var urls []string
		urls, err = postImageOperations.FindImageByPostId(l.svcCtx.DB, v.ID)
		if err != nil {
			return nil, err
		}
		// 将时间类型转换为 Unix 时间戳
		createTime := uint32(v.CreatedAt.Unix())
		// 创建新的帖子信息结构体
		newPost := &community.PostInfo{
			Id:           uint32(v.ID),
			UserId:       uint32(v.UserID),
			LikeCount:    uint32(v.LikeCount),
			Content:      v.Content,
			CreateTime:   createTime,
			ImageUrls:    urls,
			CollectCount: uint32(v.CollectionCount),
			ContentCount: uint32(v.CommentCount),
		}

		// 将新的帖子信息添加到切片中
		postInfo = append(postInfo, newPost)
	}
	return &community.CommunityLookPostByOwnResponses{PostData: postInfo}, nil
}
