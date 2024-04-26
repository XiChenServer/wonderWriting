package logic

import (
	"calligraphy/apps/community/model"
	"context"

	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeCommentLogic {
	return &LikeCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 对评论进行点赞
func (l *LikeCommentLogic) LikeComment(in *community.LikeCommentRequest) (*community.LikeCommentResponse, error) {
	// todo: add your logic here and delete this line
	err := (&model.LikeComment{}).LikeComment(l.svcCtx.DB, uint(in.CommentId), uint(in.UserId))
	if err != nil {
		l.Logger.Error("rpc 用户对帖子评论进行点赞失败，userID", in.UserId, "commentId", in.CommentId, "error", err.Error())
	}
	l.Logger.Infof("rpc 用户对评论点赞成功")
	return &community.LikeCommentResponse{}, nil
}
