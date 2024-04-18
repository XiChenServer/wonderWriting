package user

import (
	"calligraphy/apps/user/rpc/types/user"
	"context"
	"encoding/json"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFollowLogic {
	return &UserFollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFollowLogic) UserFollow(req *types.UserFollowRequest) (resp *types.UserFollowResponse, err error) {
	// todo: add your logic here and delete this line
	// todo: add your logic here and delete this line
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	_, err = l.svcCtx.UserRpc.UserFollow(l.ctx, &user.UserFollowRequest{
		UserId:  uint32(uid),
		OtherId: req.UserId,
	})
	if err != nil {
		return &types.UserFollowResponse{}, err
	}
	return &types.UserFollowResponse{}, err
	return
}
