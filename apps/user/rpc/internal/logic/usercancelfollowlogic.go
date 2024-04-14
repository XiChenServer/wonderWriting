package logic

import (
	"context"

	"calligraphy/apps/user/rpc/internal/svc"
	"calligraphy/apps/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCancelFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserCancelFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCancelFollowLogic {
	return &UserCancelFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户取消关注
func (l *UserCancelFollowLogic) UserCancelFollow(in *user.UserCancelFollowRequest) (*user.UserCancelFollowResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UserCancelFollowResponse{}, nil
}
