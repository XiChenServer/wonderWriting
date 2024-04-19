package logic

import (
	"context"

	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelLikeCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelLikeCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelLikeCommentLogic {
	return &CancelLikeCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 对评论点赞的取消
func (l *CancelLikeCommentLogic) CancelLikeComment(in *community.CancelLikeCommentRequest) (*community.CancelLikeCommentResponse, error) {
	// todo: add your logic here and delete this line

	return &community.CancelLikeCommentResponse{}, nil
}
