package group

import (
	"calligraphy/apps/group/rpc/types/group"
	"context"
	"encoding/json"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartCheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStartCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartCheckLogic {
	return &StartCheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartCheckLogic) StartCheck(req *types.StartCheckRequest) (resp *types.StartCheckResponse, err error) {
	// todo: add your logic here and delete this line
	//从jwt获取id
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.GroupRpc.StartCheck(l.ctx, &group.StartCheckRequest{UserId: uint32(uid)})
	if err != nil {
		return nil, err
	}
	return &types.StartCheckResponse{
		CheckId:         res.CheckId,
		UserId:          res.UserId,
		ContinuousDays:  res.ContinuousDays,
		CreateTime:      res.CreateTime,
		LastCheckInTime: res.LastCheckInTime,
	}, nil
}
