package logic

import (
	"calligraphy/apps/community/model"
	"context"

	"calligraphy/apps/community/rpc/internal/svc"
	"calligraphy/apps/community/rpc/types/community"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewUnreadCommentsCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewViewUnreadCommentsCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewUnreadCommentsCountLogic {
	return &ViewUnreadCommentsCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查看用户有多少未读的信息
func (l *ViewUnreadCommentsCountLogic) ViewUnreadCommentsCount(in *community.ViewUnreadCommentsCountRequest) (*community.ViewUnreadCommentsCountResponse, error) {
	// todo: add your logic here and delete this line
	count, err := (&model.UserUnreadMessages{}).UnReadMessageCount(l.svcCtx.DB, uint(in.UserId))
	if err != nil {
		return nil, err
	}
	return &community.ViewUnreadCommentsCountResponse{
		MessageCount: uint32(count),
	}, nil
}
