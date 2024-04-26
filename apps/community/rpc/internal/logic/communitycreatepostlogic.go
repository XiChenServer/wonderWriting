package logic

import (
	"calligraphy/apps/community/model"
	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type CommunityCreatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommunityCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommunityCreatePostLogic {
	return &CommunityCreatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommunityCreatePostLogic) CommunityCreatePost(in *community.CommunityCreatePostRequest) (*community.CommunityCreatePostResponse, error) {
	// todo: add your logic here and delete this line

	postOperations := model.Post{}
	res, err := postOperations.CreatePost(l.svcCtx.DB, uint(in.UserId), in.Content, in.ImageUrls)

	if err != nil {
		l.Logger.Error("rpc 用户创建帖子的时候出现了问题,err", err.Error())
		return nil, err
	}
	l.Logger.Infof("rpc 用户创建帖子成功")
	return &community.CommunityCreatePostResponse{PostId: uint32(res.ID)}, nil
}
