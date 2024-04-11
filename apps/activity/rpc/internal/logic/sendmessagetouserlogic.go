package logic

import (
	"context"

	"calligraphy/apps/activity/rpc/internal/svc"
	"calligraphy/apps/activity/rpc/types/activity"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageToUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMessageToUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageToUserLogic {
	return &SendMessageToUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMessageToUserLogic) SendMessageToUser(in *activity.SendMessageRequest) (*activity.SendMessageResponse, error) {
	// todo: add your logic here and delete this line

	return &activity.SendMessageResponse{}, nil
}
