package user

import (
	"context"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserModBackgroundLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserModBackgroundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserModBackgroundLogic {
	return &UserModBackgroundLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserModBackgroundLogic) UserModBackground() (resp *types.UserModBackgroundResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
