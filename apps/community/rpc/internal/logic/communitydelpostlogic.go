package logic

import (
	"calligraphy/apps/community/model"
	"context"

	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommunityDelPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommunityDelPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommunityDelPostLogic {
	return &CommunityDelPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommunityDelPostLogic) CommunityDelPost(in *community.CommunityDelPostRequest) (*community.CommunityDelPostResponse, error) {
	// todo: add your logic here and delete this line
	postOperations := model.Post{}
	_, err := postOperations.DeletePost(l.svcCtx.DB, in.PostId)
	if err != nil {
		return nil, err
	}
	return &community.CommunityDelPostResponse{}, nil
}
