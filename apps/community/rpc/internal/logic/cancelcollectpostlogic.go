package logic

import (
	"calligraphy/apps/community/model"
	"context"

	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelCollectPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelCollectPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelCollectPostLogic {
	return &CancelCollectPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CancelCollectPostLogic) CancelCollectPost(in *community.CommunityCancelCollectPostRequest) (*community.CommunityCancelCollectPostResponse, error) {
	// todo: add your logic here and delete this line
	Operations := &model.Collect{}
	err := Operations.CancelCollectPost(l.svcCtx.DB, uint(in.UserId), uint(in.PostId))
	if err != nil {
		return nil, err
	}
	return &community.CommunityCancelCollectPostResponse{}, nil
}
