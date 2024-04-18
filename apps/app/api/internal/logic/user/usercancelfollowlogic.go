package user

import (
	"calligraphy/apps/user/rpc/types/user"
	"context"
	"encoding/json"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCancelFollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCancelFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCancelFollowLogic {
	return &UserCancelFollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCancelFollowLogic) UserCancelFollow(req *types.UserCancelFollowRequest) (resp *types.UserCancelFollowResponse, err error) {
	// todo: add your logic here and delete this line
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	_, err = l.svcCtx.UserRpc.UserCancelFollow(l.ctx, &user.UserCancelFollowRequest{
		UserId:  uint32(uid),
		OtherId: req.UserId,
	})
	if err != nil {
		return &types.UserCancelFollowResponse{}, err
	}
	return &types.UserCancelFollowResponse{}, err
}
