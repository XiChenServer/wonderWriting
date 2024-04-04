package logic

import (
	"context"

	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"

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

	return &community.CommunityLookPostByOwnResponses{}, nil
}
