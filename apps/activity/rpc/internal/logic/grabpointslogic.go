package logic

import (
	"context"

	"calligraphy/apps/activity/rpc/internal/svc"
	"calligraphy/apps/activity/rpc/types/activity"

	"github.com/zeromicro/go-zero/core/logx"
)

type GrabPointsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGrabPointsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrabPointsLogic {
	return &GrabPointsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GrabPointsLogic) GrabPoints(in *activity.GrabPointsRequest) (*activity.GrabPointsResponse, error) {
	// todo: add your logic here and delete this line

	return &activity.GrabPointsResponse{}, nil
}
