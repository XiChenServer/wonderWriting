package user

import (
	"context"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserModPwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserModPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserModPwdLogic {
	return &UserModPwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserModPwdLogic) UserModPwd(req *types.UserModPwdRequset) (resp *types.UserModPwdResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
