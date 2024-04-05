package logic

import (
	"context"

	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"

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
	// todo: add your logic here and delete this line

	return &community.CommunityLookAllPostsResponse{}, nil
}
