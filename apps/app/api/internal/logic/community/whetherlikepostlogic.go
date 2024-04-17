package community

import (
	"calligraphy/apps/community/rpc/types/community"
	"context"
	"encoding/json"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WhetherLikePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWhetherLikePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WhetherLikePostLogic {
	return &WhetherLikePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WhetherLikePostLogic) WhetherLikePost(req *types.WhetherLikePostRequest) (resp *types.WhetherLikePostResponse, err error) {
	// todo: add your logic here and delete this line
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	_, err = l.svcCtx.CommunityRpc.WhetherLikePost(l.ctx, &community.WhetherLikePostRequest{
		UserId: uint32(uid),
		PostId: req.OtherId,
	})
	if err != nil {
		return nil, err
	}

	return &types.WhetherLikePostResponse{}, nil

}
