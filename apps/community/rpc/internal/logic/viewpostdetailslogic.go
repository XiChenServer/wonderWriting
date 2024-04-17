package logic

import (
	"context"

	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"

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

	return &community.ViewPostDetailsResponse{}, nil
}
