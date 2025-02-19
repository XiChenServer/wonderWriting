package community

import (
	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"
	"calligraphy/apps/community/rpc/types/community"
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelCollectPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelCollectPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelCollectPostLogic {
	return &CancelCollectPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelCollectPostLogic) CancelCollectPost(req *types.CancelCollectPostRequest) (resp *types.CancelCollectPostResponse, err error) {
	// todo: add your logic here and delete this line
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	_, err = l.svcCtx.CommunityRpc.CancelCollectPost(l.ctx, &community.CommunityCancelCollectPostRequest{PostId: uint32(req.PostId), UserId: uint32(uid)})
	if err != nil {
		return &types.CancelCollectPostResponse{}, err
	}
	return &types.CancelCollectPostResponse{}, nil
}
