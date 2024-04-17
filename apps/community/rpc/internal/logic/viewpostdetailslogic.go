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

type ViewPostDetailsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewViewPostDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewPostDetailsLogic {
	return &ViewPostDetailsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func (l *ViewPostDetailsLogic) ViewPostDetails(in *community.ViewPostDetailsRequest) (*community.ViewPostDetailsResponse, error) {
	// todo: add your logic here and delete this line
	resPost, err := (&model.Post{}).LookPostByPostId(l.svcCtx.DB, uint(in.PostId))
	if err != nil {
		return nil, err
	}

	resUser, err := (&userModel.User{}).FindOne(l.svcCtx.DB, resPost.UserID)
	if err != nil {
		return nil, err
	}

	ansUser := &community.UserSimpleInfo{
		Id:          uint32(resUser.UserID),
		NickName:    resUser.Nickname,
		Account:     resUser.Account,
		AvatarImage: qiniu.ImgUrl + resUser.AvatarBackground,
	}

	urls, err := (&model.PostImage{}).FindImageByPostId(l.svcCtx.DB, resPost.ID)
	if err != nil {
		return nil, err
	}
	ansPost := &community.PostInfo{
		Id:           uint32(resPost.ID),
		UserId:       uint32(resPost.UserID),
		LikeCount:    uint32(resPost.LikeCount),
		Content:      resPost.Content,
		ImageUrls:    urls,
		CreateTime:   uint32(resPost.CreatedAt.Unix()),
		ContentCount: uint32(resPost.CommentCount),
		CollectCount: uint32(resPost.CollectionCount),
	}
	fmt.Println("2123")
	ansPost.UserInfo = ansUser
	fmt.Println("2123")
	var ansStatus = &community.StatusWithPost{
		WhetherLike:    false,
		WhetherCollect: false,
		WhetherFollow:  false,
	}

	// 检查用户关注状态
	err = (&userModel.Follow{}).WhetherFollowPeople(l.svcCtx.DB, resUser.UserID, uint(in.UserId))
	if err == nil {
		ansStatus.WhetherFollow = true
	}

	// 检查用户点赞状态
	err = (&model.Like{}).WhetherLikedPost(l.svcCtx.DB, resPost.ID, uint(in.UserId))
	if err == nil {
		ansStatus.WhetherLike = true
	}

	// 检查用户收藏状态
	err = (&model.Collect{}).WhetherCollectPost(l.svcCtx.DB, resPost.ID, uint(in.UserId))
	if err == nil {
		ansStatus.WhetherCollect = true
	}

	return &community.ViewPostDetailsResponse{
		PostData:      ansPost,
		RelatedStatus: ansStatus,
	}, nil
}
