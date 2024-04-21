package community

import (
	"calligraphy/apps/community/rpc/types/community"
	"context"
	"encoding/json"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReplyCommunityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReplyCommunityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReplyCommunityLogic {
	return &ReplyCommunityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReplyCommunityLogic) ReplyCommunity(req *types.ReplyCommunityRequest) (resp *types.ReplyCommentResponse, err error) {
	// todo: add your logic here and delete this line
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.CommunityRpc.ReplyComment(l.ctx, &community.ReplyCommunityRequest{
		CommunityId:       req.CommunityId,
		Content:           req.Content,
		UserId:            uint32(uid),
		ReplyUserNickName: req.ReplyUserNickName,
		ReplyUserId:       req.ReplyUserId,
		PostId:            req.PostId,
	})
	return &types.ReplyCommentResponse{
		ReplyCommunityId: res.ReplyCommunityId,
	}, nil
}
