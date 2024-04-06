package logic

import (
	"calligraphy/apps/community/model"
	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"
	userModel "calligraphy/apps/user/model"
	"calligraphy/pkg/qiniu"
	"context"
	"fmt"
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

func (l *CommunityLookAllPostsLogic) CommunityLookAllPosts(in *community.CommunityLookAllPostsRequest) (*community.CommunityLookAllPostsResponse, error) {
	// 创建 Post 和 PostImage 操作实例
	postOperations := model.Post{}
	postImageOperations := model.PostImage{}
	// 查询所有帖子信息
	res, err := postOperations.LookAllPosts(l.svcCtx.DB)
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

		var User userModel.User
		err = l.svcCtx.DB.Where("user_id = ?", v.UserID).First(&User).Error
		if err != nil {
			fmt.Println(err, v.UserID)
			return nil, err
		}
		var userInfo = community.UserSimpleInfo{
			Id:          uint32(User.UserID),
			NickName:    User.Nickname,
			Account:     User.Account,
			AvatarImage: qiniu.ImgUrl + User.AvatarBackground,
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
			UserInfo:     &userInfo,
		}
		// 将新的帖子信息添加到切片中
		postInfo = append(postInfo, newPost)
	}

	// 构建并返回帖子信息响应
	return &community.CommunityLookAllPostsResponse{
		PostData: postInfo,
	}, nil
}
