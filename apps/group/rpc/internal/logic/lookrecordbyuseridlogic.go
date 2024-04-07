package logic

import (
	"context"

	"calligraphy/apps/group/rpc/internal/svc"
	"calligraphy/apps/group/rpc/types/group"

	"github.com/zeromicro/go-zero/core/logx"
)

type LookRecordByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLookRecordByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LookRecordByUserIdLogic {
	return &LookRecordByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查看某人的书法记录
func (l *LookRecordByUserIdLogic) LookRecordByUserId(in *group.LookRecordByUserIdRequest) (*group.LookRecordByUserIdResponse, error) {
	// todo: add your logic here and delete this line

	return &group.LookRecordByUserIdResponse{}, nil
}
