package logic

import (
	"calligraphy/apps/community/model"
	"context"

	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentPostLogic {
	return &CommentPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 定义评论服务
func (l *CommentPostLogic) CommentPost(in *community.CommunityContentPostRequest) (*community.CommunityContentPostResponse, error) {
	// todo: add your logic here and delete this line
	Operations := &model.Comment{}
	res, err := Operations.CommentPost(l.svcCtx.DB, uint(in.PostId), uint(in.UserId), in.Content)
	if err != nil {
		return nil, err
	}
	return &community.CommunityContentPostResponse{ContentId: uint32(res.ID)}, nil
}
