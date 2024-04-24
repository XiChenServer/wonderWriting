package community

import (
	"calligraphy/apps/community/rpc/types/community"
	"context"
	"encoding/json"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewUnreadCommentCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewViewUnreadCommentCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewUnreadCommentCountLogic {
	return &ViewUnreadCommentCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ViewUnreadCommentCountLogic) ViewUnreadCommentCount(req *types.ViewUnreadCommentCountRequest) (resp *types.ViewUnreadCommentCountResponse, err error) {
	// todo: add your logic here and delete this line
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.CommunityRpc.ViewUnreadCommentsCount(l.ctx, &community.ViewUnreadCommentsCountRequest{UserId: uint32(uid)})
	if err != nil {
		return nil, err
	}
	return &types.ViewUnreadCommentCountResponse{UnreadCommentCount: uint64(res.MessageCount)}, nil
}
