package community

import (
	"calligraphy/apps/community/rpc/types/community"
	"context"
	"encoding/json"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentPostLogic {
	return &CommentPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentPostLogic) CommentPost(req *types.CommentPostRequest) (resp *types.CommentPostResponse, err error) {
	// todo: add your logic here and delete this line
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.CommunityRpc.CommentPost(l.ctx, &community.CommunityContentPostRequest{
		PostId:  uint32(req.PostId),
		UserId:  uint32(uid),
		Content: req.Content,
	})
	if err != nil {
		return &types.CommentPostResponse{}, err
	}
	return &types.CommentPostResponse{CommentId: uint(res.ContentId)}, nil
}
