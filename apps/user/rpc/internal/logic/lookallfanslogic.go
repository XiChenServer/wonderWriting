package logic

import (
	"context"

	"calligraphy/apps/user/rpc/internal/svc"
	"calligraphy/apps/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LookAllFansLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLookAllFansLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookAllFansLogic {
	return &LookAllFansLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户查看自己的粉丝
func (l *LookAllFansLogic) LookAllFans(in *user.LookAllFansRequest) (*user.LookAllFansResponse, error) {
	// todo: add your logic here and delete this line

	return &user.LookAllFansResponse{}, nil
}
