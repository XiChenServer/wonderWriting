package community

import (
	"calligraphy/apps/community/rpc/types/community"
	"context"
	"encoding/json"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelCommentPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelCommentPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelCommentPostLogic {
	return &CancelCommentPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelCommentPostLogic) CancelCommentPost(req *types.CancelCommentPostRequest) (resp *types.CancelCommentPostResponse, err error) {
	// todo: add your logic here and delete this line
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()

	_, err = l.svcCtx.CommunityRpc.CancelCommentPost(l.ctx, &community.CommunityCancelContentPostRequest{PostId: uint32(req.PostId), UserId: uint32(uid), ContentId: uint32(req.CommentId)})
	if err != nil {
		return &types.CancelCommentPostResponse{}, err
	}
	return &types.CancelCommentPostResponse{}, nil
}
