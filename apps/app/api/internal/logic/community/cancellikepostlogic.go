package community

import (
	"calligraphy/apps/community/rpc/types/community"
	"context"
	"encoding/json"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelLikePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelLikePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelLikePostLogic {
	return &CancelLikePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelLikePostLogic) CancelLikePost(req *types.CancelLikePostRequest) (resp *types.CancelLikePostResponse, err error) {
	// todo: add your logic here and delete this line
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()

	_, err = l.svcCtx.CommunityRpc.CancelLikePost(l.ctx, &community.CommunityCancelLikePostRequest{
		LikeId: uint32(req.LikeId),
		UserId: uint32(uid),
		PostId: uint32(req.PostId),
	})
	if err != nil {
		return nil, err
	}
	return &types.CancelLikePostResponse{}, nil
}
