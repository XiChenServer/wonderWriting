package group

import (
	"calligraphy/apps/group/rpc/types/group"
	"context"
	"encoding/json"

	"calligraphy/apps/app/api/internal/svc"
	"calligraphy/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckPunchCardLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckPunchCardLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckPunchCardLogic {
	return &CheckPunchCardLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckPunchCardLogic) CheckPunchCard(req *types.CheckPunchCardModelRequest) (resp *types.CheckPunchCardModelResponse, err error) {
	// todo: add your logic here and delete this line
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.GroupRpc.CheckPunchCardModel(l.ctx, &group.CheckPunchCardModelRequest{UserId: uint32(uid)})
	if err != nil {
		return nil, err
	}

	return &types.CheckPunchCardModelResponse{Data: res.Data}, nil
}
