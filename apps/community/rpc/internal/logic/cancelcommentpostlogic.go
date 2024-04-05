package logic

import (
	"calligraphy/apps/community/model"
	"context"

	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelCommentPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelCommentPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelCommentPostLogic {
	return &CancelCommentPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CancelCommentPostLogic) CancelCommentPost(in *community.CommunityCancelContentPostRequest) (*community.CommunityCancelContentPostResponse, error) {
	// todo: add your logic here and delete this line
	Operations := &model.Comment{}
	err := Operations.CancelCommentPost(l.svcCtx.DB, uint(in.ContentId), uint(in.PostId))
	if err != nil {
		return nil, err
	}
	return &community.CommunityCancelContentPostResponse{}, nil
}
