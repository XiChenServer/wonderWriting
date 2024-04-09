package activity

import (
	"context"
	"fmt"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GrabPointsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGrabPointsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GrabPointsLogic {
	return &GrabPointsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GrabPointsLogic) GrabPoints(req *types.GrabPointsRequest) (resp *types.GrabPointsResponse, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("123")
	data := "zhangSan"
	if err := l.svcCtx.KqPusherClient.Push(data); err != nil {
		logx.Errorf("KqPusherClient Push Error , err :%v", err)
	}
	return
}
