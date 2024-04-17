package user

import (
	"calligraphy/apps/user/rpc/types/user"
	"context"
	"encoding/json"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WhetherFollowUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWhetherFollowUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WhetherFollowUserLogic {
	return &WhetherFollowUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WhetherFollowUserLogic) WhetherFollowUser(req *types.WhetherFollowUserRequest) (resp *types.WhetherFollowUserResponse, err error) {
	// todo: add your logic here and delete this line
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	_, err = l.svcCtx.UserRpc.WhetherFollowUser(l.ctx, &user.WhetherFollowUserRequest{
		UserId:  uint32(uid),
		OtherId: req.OtherId,
	})
	if err != nil {
		return nil, err
	}

	return &types.WhetherFollowUserResponse{}, nil
}
