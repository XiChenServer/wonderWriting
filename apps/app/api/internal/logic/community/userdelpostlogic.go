package community

import (
	"calligraphy/apps/community/rpc/types/community"
	"context"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDelPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDelPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDelPostLogic {
	return &UserDelPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDelPostLogic) UserDelPost(req *types.PostDelRequest) (resp *types.PostDelResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.CommunityRpc.CommunityDelPost(l.ctx, &community.CommunityDelPostRequest{PostId: uint32(req.PostId)})
	if err != nil {
		return nil, err
	}
	return &types.PostDelResponse{}, nil
}
