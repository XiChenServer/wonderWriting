package logic

import (
	"context"

	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"

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

	return &community.CommunityCreatePostResponse{}, nil
}
