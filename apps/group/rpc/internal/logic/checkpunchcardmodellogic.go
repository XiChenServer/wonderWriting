package logic

import (
	groupModel "calligraphy/apps/group/model"
	"context"

	"calligraphy/apps/group/rpc/internal/svc"
	"calligraphy/apps/group/rpc/types/group"

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
func (l *CheckPunchCardModelLogic) CheckPunchCardModel(in *group.CheckPunchCardModelRequest) (*group.CheckPunchCardModelResponse, error) {
	// todo: add your logic here and delete this line
	res, err := (&groupModel.CheckIn{}).IsCheckInOpen(l.svcCtx.DB, uint(in.UserId))
	if err != nil {
		return nil, err
	}
	return &group.CheckPunchCardModelResponse{
		Data: res,
	}, nil
}
