package user

import (
	"context"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserModAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserModAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserModAvatarLogic {
	return &UserModAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserModAvatarLogic) UserModAvatar() (resp *types.UserModAvatarResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
