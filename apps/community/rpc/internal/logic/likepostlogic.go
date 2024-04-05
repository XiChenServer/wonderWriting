package logic

import (
	"calligraphy/apps/community/model"
	"context"

	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikePostLogic {
	return &LikePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 定义点赞服务
func (l *LikePostLogic) LikePost(in *community.CommunityLikePostRequest) (*community.CommunityLikePostResponse, error) {
	// todo: add your logic here and delete this line
	Operations := &model.Like{}
	res, err := Operations.LikePost(l.svcCtx.DB, uint(in.PostId), uint(in.UserId))
	if err != nil {
		return nil, err
	}
	return &community.CommunityLikePostResponse{LikeId: uint32(res.ID)}, nil
}
