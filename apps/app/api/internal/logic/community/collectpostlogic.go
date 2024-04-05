package community

import (
	"calligraphy/apps/community/rpc/types/community"
	"context"
	"encoding/json"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCollectPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectPostLogic {
	return &CollectPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CollectPostLogic) CollectPost(req *types.CollectPostRequest) (resp *types.CollectPostResponse, err error) {
	// todo: add your logic here and delete this line
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.CommunityRpc.CollectPost(l.ctx, &community.CommunityCollectPostRequest{
		PostId: uint32(req.PostId),
		UserId: uint32(uid),
	})
	if err != nil {
		return nil, err
	}
	return &types.CollectPostResponse{CollectId: uint(res.CollectId)}, nil
}
