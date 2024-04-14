package logic

import (
	"context"

	"calligraphy/apps/user/rpc/internal/svc"
	"calligraphy/apps/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFollowLogic {
	return &UserFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户关注
func (l *UserFollowLogic) UserFollow(in *user.UserFollowRequest) (*user.UserFollowResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UserFollowResponse{}, nil
}
