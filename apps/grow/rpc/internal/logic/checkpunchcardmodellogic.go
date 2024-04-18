package logic

import (
	groupModel "calligraphy/apps/grow/model"
	"context"

	"calligraphy/apps/grow/rpc/internal/svc"
	"calligraphy/apps/grow/rpc/types/grow"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckPunchCardModelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckPunchCardModelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckPunchCardModelLogic {
	return &CheckPunchCardModelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 检查打卡模式是否开启
func (l *CheckPunchCardModelLogic) CheckPunchCardModel(in *grow.CheckPunchCardModelRequest) (*grow.CheckPunchCardModelResponse, error) {
	// todo: add your logic here and delete this line
	res, err := (&groupModel.CheckIn{}).IsCheckInOpen(l.svcCtx.DB, uint(in.UserId))
	if err != nil {
		return nil, err
	}
	return &grow.CheckPunchCardModelResponse{
		Data: res,
	}, nil
}
