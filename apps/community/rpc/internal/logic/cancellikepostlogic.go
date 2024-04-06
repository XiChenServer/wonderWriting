package logic

import (
	"calligraphy/apps/community/model"
	"context"

	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelLikePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelLikePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelLikePostLogic {
	return &CancelLikePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CancelLikePostLogic) CancelLikePost(in *community.CommunityCancelLikePostRequest) (*community.CommunityCancelLikePostResponse, error) {
	// todo: add your logic here and delete this line
	Operations := &model.Like{}
	err := Operations.CancelLikePost(l.svcCtx.DB, uint(in.LikeId), uint(in.PostId), uint(in.UserId))
	if err != nil {
		return nil, err
	}
	return &community.CommunityCancelLikePostResponse{}, nil
}
