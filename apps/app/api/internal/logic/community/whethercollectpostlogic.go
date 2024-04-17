package community

import (
	"calligraphy/apps/community/rpc/types/community"
	"context"
	"encoding/json"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WhetherCollectPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWhetherCollectPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WhetherCollectPostLogic {
	return &WhetherCollectPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WhetherCollectPostLogic) WhetherCollectPost(req *types.WhetherCollectPostRequest) (resp *types.WhetherCollectPostResponse, err error) {
	// todo: add your logic here and delete this line
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	_, err = l.svcCtx.CommunityRpc.WhetherCollectPost(l.ctx, &community.WhetherCollectPostRequest{
		UserId: uint32(uid),
		PostId: req.OtherId,
	})
	if err != nil {
		return nil, err
	}

	return &types.WhetherCollectPostResponse{}, nil

}
