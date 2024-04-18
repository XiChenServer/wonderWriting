package logic

import (
	groupModel "calligraphy/apps/grow/model"
	"calligraphy/apps/grow/rpc/internal/svc"
	"calligraphy/apps/grow/rpc/types/grow"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartCheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStartCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartCheckLogic {
	return &StartCheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 开启签到
func (l *StartCheckLogic) StartCheck(in *grow.StartCheckRequest) (*grow.StartCheckResponse, error) {
	// todo: add your logic here and delete this line
	//在model层进行创建
	res, err := (&groupModel.CheckIn{}).CreateCheckIn(l.svcCtx.DB, uint(in.UserId))
	if err != nil {
		return nil, err
	}
	//将数据返回
	return &grow.StartCheckResponse{
		CheckId:         uint32(res.ID),
		UserId:          uint32(res.UserID),
		ContinuousDays:  res.ContinuousDays,
		CreateTime:      int32(res.CreatedAt.Unix()),       // 将时间戳转换为 int32 表示
		LastCheckInTime: int32(res.LastCheckInTime.Unix()), // 将时间戳转换为 int32 表示
	}, nil
}
