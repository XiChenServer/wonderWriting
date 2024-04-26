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
	res, err := Operations.CreateComment(l.svcCtx.DB, uint(in.PostId), uint(in.UserId), in.Content)
	if err != nil {
		l.Logger.Error("rpc 用户在评论的时候，数据库操作出现问题，err:", err.Error())
		return nil, err
	}
	l.Logger.Infof("rpc 用户对帖子创建评论成功")
	return &community.CommunityContentPostResponse{ContentId: uint32(res.ID)}, nil
}
