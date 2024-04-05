package logic

import (
	"context"

	"calligraphy/apps/community/api/internal/svc"
	"calligraphy/apps/community/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get() (resp *types.GetResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
