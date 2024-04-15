package logic

import (
	"context"

	"calligraphy/apps/user/rpc/internal/svc"
	"calligraphy/apps/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LookAllFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLookAllFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookAllFollowLogic {
	return &LookAllFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户查看自己的关注
func (l *LookAllFollowLogic) LookAllFollow(in *user.LookAllFollowRequest) (*user.LookAllFollowResponse, error) {
	// todo: add your logic here and delete this line

	return &user.LookAllFollowResponse{}, nil
}
