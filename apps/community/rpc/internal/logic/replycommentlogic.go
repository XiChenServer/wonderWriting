package logic

import (
	"calligraphy/apps/community/model"
	"context"

	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReplyCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReplyCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReplyCommentLogic {
	return &ReplyCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 回复评论
func (l *ReplyCommentLogic) ReplyComment(in *community.ReplyCommunityRequest) (*community.ReplyCommunityResponse, error) {
	// todo: add your logic here and delete this line
	_, err := (&model.ReplyComment{}).ReplyComment(l.svcCtx.DB, uint(in.CommunityId), uint(in.UserId), uint(in.ReplyUserId), uint(in.PostId),
		in.ReplyUserNickName, in.Content)
	if err != nil {
		l.Error("rpc 用户在回复评论的时候, 数据库操作出现了问题，err", err.Error())
		return nil, err
	}
	l.Infof("rpc 用户回复评论成功")
	return &community.ReplyCommunityResponse{}, nil
}
