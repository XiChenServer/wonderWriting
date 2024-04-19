package logic

import (
	"calligraphy/apps/community/model"
	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCollectPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectPostLogic {
	return &CollectPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CollectPostLogic) CollectPost(in *community.CommunityCollectPostRequest) (*community.CommunityCollectPostResponse, error) {
	// todo: add your logic here and delete this line
	Operations := &model.Collect{}
	res, err := Operations.CollectPost(l.svcCtx.DB, uint(in.PostId), uint(in.UserId))
	if err != nil {
		return nil, err
	}
	return &community.CommunityCollectPostResponse{CollectId: uint32(res.ID)}, nil
}
